FROM golang:1.14 as builder

COPY /app /app
WORKDIR /app

RUN make build

ENTRYPOINT make launch
