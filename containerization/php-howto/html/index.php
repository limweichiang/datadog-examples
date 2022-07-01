<?php
  require __DIR__ . '/vendor/autoload.php';

  // load Monolog library
  use Monolog\Logger;
  use Monolog\Handler\StreamHandler;
  use Monolog\Formatter\JsonFormatter;

  // create a log channel
  $logger = new Logger('default');

  // create a Json formatter
  $formatter = new JsonFormatter();

  // create a handler
  $stream = new StreamHandler("php://stdout");
  $stream->setFormatter($formatter);

  // bind
  $logger->pushHandler($stream);

  // Datadog APM tracing - automatically append the identifiers to all log messages
  $logger->pushProcessor(function ($record) {
    // Extract trace context and inject into JSON log
    $context = \DDTrace\current_context();
    $record->extra['dd'] = [
      'trace_id' => $context['trace_id'],
      'span_id'  => $context['span_id']
    ];
    return $record;
  });

  // an example
  $logger->info('Hello World!');

  phpinfo();
?>
