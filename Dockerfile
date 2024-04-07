FROM golang:1.22-alpine AS builder

COPY . /github.com/gogapopp/trainee-assignment/source/
WORKDIR /github.com/gogapopp/trainee-assignment/source/

RUN go mod download
RUN go build -o ./bin/banner cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/gogapopp/trainee-assignment/source/bin/banner .

CMD ["./banner"]