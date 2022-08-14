# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

ADD . /app
WORKDIR /app/cmd

RUN go mod download

RUN go build -o /goscat 

CMD ["/goscat"]

# How to build:
#   docker build --network=host -t goscat:latest -f ./build/Dockerfile .