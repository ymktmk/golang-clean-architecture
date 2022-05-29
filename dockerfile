FROM --platform=linux/x86_64 golang:1.18

WORKDIR /work

RUN go install github.com/cosmtrek/air@v1.29.0

EXPOSE 9000

CMD ["air", "-c", ".air.toml"]