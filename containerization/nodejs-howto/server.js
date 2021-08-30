'use strict';

const express = require('express');
const winston = require('winston');

const logger = winston.createLogger({
  level: 'info',
  exitOnError: false,
  format: winston.format.json(),
  transports: [
    new winston.transports.Console(),
  ],
});

module.exports = logger;

// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

// App
const app = express();
app.get('/', (req, res) => {
  logger.log('info', 'Received request, sending back "Hello World"');
  res.send('Hello World');
});

app.listen(PORT, HOST);
logger.log('info', 'Running on http://${HOST}:${PORT}');
