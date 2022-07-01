# PHP Application Containerization with Datadog Howto
## Overview

This example shows how to inject a Datadog APM tracer into a PHP "Hello World!" application container.

This example (especially the Dockerfile) can be used as a reference for other development languages as well.

## Build steps

Download all files within this folder level (php-howto) and run the following command within the folder.
```
$ docker build -t <YOUR_BUILD_TAG> .
```

## Test Steps

Run the containerized test application with the following command.
```
docker run -d -p 8080:80 --name php-howto-test <YOUR_BUILD_TAG>
```

Point your browser at `http://localhost:8080/` and you should see a printout of `php_info()`.

Check to see APM injected log traces with the following command.
```
docker logs php-howto-test
```

To set up full tracing, you will need to run the following command:
```
docker run -d -p 8080:80 --network app-bridge \
  && -e DD_AGENT_HOST=datadog-agent \
  && -e DD_ENV=<ENVIRONMENT_NAME> \
  && -e DD_SERVICE=<SERVICE_NAME> \
  && -e DD_VERSION=<VERSION_NUMBER> \
  && --name php-howto-test <YOUR_BUILD_TAG>
```
This assumes you have set up the Datadog Docker agent with a network name of `app-bridge` following the instructions for Tracing from other containers(https://docs.datadoghq.com/containers/docker/apm/?tab=java&tabs=standard#tracing-from-other-containers), and have enabled logging for Docker containers following the instructions at [Docker Log Collection - Container Installation](https://docs.datadoghq.com/containers/docker/log/?tabs=dockerfile#installation)