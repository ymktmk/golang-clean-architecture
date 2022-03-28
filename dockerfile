FROM golang:latest

RUN mkdir /go/src/work

WORKDIR /go/src/work

COPY . /go/src/work

EXPOSE 9000