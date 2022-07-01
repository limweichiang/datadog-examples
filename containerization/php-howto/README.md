# PHP Application Containerization with Datadog Howto
## Overview

This example shows how to inject a Datadog APM tracer into a PHP "Hello World!" application container.

This example (especially the Dockerfile) can be used as a reference for other development languages as well.

## Build steps

Download all files within this folder level (php-howto) and run the following command within the folder.
```
$ docker build -t <YOUR_BUILD_TAG_NAME> .
```

## Test Steps

Run the containerized test application with the following command.
```
docker run -d -p 8080:80 --name php-howto-test <YOUR_BUILD_TAG_NAME>
```

Point your browser at `http://localhost:8080/` and you should see a printout of `php_info()`.

Check to see APM injected log traces with the following command.
```
docker logs php-howto-test
```