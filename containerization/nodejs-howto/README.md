# Node.js Application Containerization with Datadog Howto
## Overview

This example shows how to inject a Datadog APM tracer into a Node.js "Hello World!" application container. This scenario builds upon the example that is at the [Dockerizing a Node.js web app](https://nodejs.org/en/docs/guides/nodejs-docker-webapp/) example at the Node.js website.

This example (especially the Dockerfile) can be used as a reference for other development languages as well.

## Steps

Download all files within this folder level (nodejs-howto) and run the following command within the folder
```
$ docker build -t <your-tag-name> .
```
