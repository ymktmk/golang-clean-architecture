FROM --platform=linux/x86_64 golang:1.18

RUN mkdir /go/src/work

WORKDIR /go/src/work

COPY ./ /go/src/work

EXPOSE 9000