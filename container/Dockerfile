FROM golang:1.21

RUN set -x && \
  apt-get update -y && \
  apt-get install -y dieharder && \
  apt-get install -y pv

WORKDIR /app
