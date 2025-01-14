FROM golang:1.24rc1-alpine3.21 AS builder

WORKDIR /src

COPY ./src .

RUN go mod download

RUN go build -o fiesta-box-server

FROM alpine:3.20.3 AS compiled

WORKDIR /root

COPY --from=builder /src/fiesta-box-server .

EXPOSE 8080

ENTRYPOINT ["./fiesta-box-server"]