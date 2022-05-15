FROM golang:1.18.2-alpine

RUN mkdir /app
WORKDIR /app

COPY go.mod /app
COPY go.sum /app
RUN go mod download
