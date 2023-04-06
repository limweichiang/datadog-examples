package main

import (
    // Standard libraries
    "fmt"
    "net/http"
    "io"
    "encoding/json"
    "context"

    // Additional frameworks / libraries
    "github.com/gin-gonic/gin"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"

    // Datadog default APM tracer
    "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
    "gopkg.in/DataDog/dd-trace-go.v1/profiler"

    // Datadog APM-supported frameworks/libraries
    gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
    httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

func main() {
    // Configure log format for Unix timestamps with millisecond granularity
    // for log correlation precision in trace.
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

    // Start Datadog Tracing
    tracer.Start()
    defer tracer.Stop()

    // If you need to start logging from within main, explicitly starting a
    // span and extracting trace & span IDs is necessary to inject these
    // values into the log. If you don't, the span is implicitly started
    // anyway, except there's no explicit span instance to extract from.
    span,_ := tracer.StartSpanFromContext(context.Background(),"")
    defer span.Finish()
    span_id := span.Context().SpanID()
    trace_id := span.Context().TraceID()
    log.Info().
        RawJSON("dd.trace_id", []byte(fmt.Sprint(trace_id))).
        RawJSON("dd.span_id", []byte(fmt.Sprint(span_id))).
        Msg("This is a log generated from within main()")

    // Start Datadog Continuous Profiler. Optional.
    err := profiler.Start(
        profiler.WithProfileTypes(
          profiler.CPUProfile,
          profiler.HeapProfile,
          // The profiles below are disabled by default to keep overhead
          // low, but can be enabled as needed.
          // profiler.BlockProfile,
          // profiler.MutexProfile,
          // profiler.GoroutineProfile,
        ),
    )
    if err != nil {
        log.Error().Err(err)
    }
    defer profiler.Stop()
    
    // Init Gin instance, and wrap with Datadog tracer
    r := gin.New()
    r.Use(gintrace.Middleware(""))

    // Default ping route
    r.GET("/ping", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
        "message": "pong",
      })
    })

    // Create joke route
    r.GET("/joke", func(c *gin.Context) {
        // Extract current span from context, for injecting context into logs.
        span, _ := tracer.SpanFromContext(c.Request.Context())
        span_id := span.Context().SpanID()
        trace_id := span.Context().TraceID()

        // Create URL string, used for HTTP request and service name
        url := "https://official-joke-api.appspot.com/random_joke"

        // Create http client, and wrap with Datadog tracer
        client := &http.Client{}
        httptrace.WrapClient(client, httptrace.RTWithServiceName(url))

        // Setup and execute GET request to the url; Context must be passed to link http client request span to the trace.
        log.Info().
            RawJSON("dd.trace_id", []byte(fmt.Sprint(trace_id))).
            RawJSON("dd.span_id", []byte(fmt.Sprint(span_id))).
            Msg("Making GET request to " + url)
        req, _ := http.NewRequestWithContext(c.Request.Context(), "GET", url, nil)
        resp, err := client.Do(req)

        if err != nil {
            log.Error().
                RawJSON("dd.trace_id", []byte(fmt.Sprint(trace_id))).
                RawJSON("dd.span_id", []byte(fmt.Sprint(span_id))).
                Err(err)
        }
        defer resp.Body.Close()

        // Process the response body
        log.Debug().
            RawJSON("dd.trace_id", []byte(fmt.Sprint(trace_id))).
            RawJSON("dd.span_id", []byte(fmt.Sprint(span_id))).
            Msg("Processing HTTP response body")

		body, err := io.ReadAll(resp.Body)

        // Check for read errors
        if err != nil {
            log.Error().
                RawJSON("dd.trace_id", []byte(fmt.Sprint(trace_id))).
                RawJSON("dd.span_id", []byte(fmt.Sprint(span_id))).
                Err(err)
            
            // Set return for calling client
            c.JSON(http.StatusInternalServerError, string(body))
        }
        
        // Check for valid JSON formatted response
        if !json.Valid(body){
            log.Error().
                RawJSON("dd.trace_id", []byte(fmt.Sprint(trace_id))).
                RawJSON("dd.span_id", []byte(fmt.Sprint(span_id))).
                Msg("Did not receive valid JSON response from endpoint")
            
            // Set return for calling client
            c.JSON(http.StatusInternalServerError, string(body))
        }

        log.Info().
            RawJSON("dd.trace_id", []byte(fmt.Sprint(trace_id))).
            RawJSON("dd.span_id", []byte(fmt.Sprint(span_id))).
            RawJSON("response", body).
            Msg("Received valid JSON from endpoint")
        
        // Set return for calling client
        c.JSON(http.StatusOK, string(body))
    })

    r.Run()
}
