FROM php:8.1-apache

# Copying in the files we'll need
COPY html /var/www/html
COPY tmp /tmp

WORKDIR /tmp

## Full installation: APM + Application Security + Profiling
RUN curl -LO https://github.com/DataDog/dd-trace-php/releases/latest/download/datadog-setup.php \
    && php datadog-setup.php --php-bin=all --enable-appsec --enable-profiling

## APM only
# RUN curl -LO https://github.com/DataDog/dd-trace-php/releases/latest/download/datadog-setup.php \
#    && php datadog-setup.php --php-bin=all

## APM with Application Security
# RUN curl -LO https://github.com/DataDog/dd-trace-php/releases/latest/download/datadog-setup.php \
#    && php datadog-setup.php --php-bin=all --enable-appsec

## APM with Profiling
# RUN curl -LO https://github.com/DataDog/dd-trace-php/releases/latest/download/datadog-setup.php \
#     && php datadog-setup.php --php-bin=all --enable-profiling

# Install Composer and Monolog logging framework
RUN /tmp/composer-setup.sh
WORKDIR /var/www/html
RUN apt update \
    && apt install -y git unzip \
    && php /tmp/composer.phar require "monolog/monolog"

# Expose port 80
EXPOSE 80

