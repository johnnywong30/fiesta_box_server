FROM golang:1.24rc1-alpine3.21

WORKDIR /src

RUN go install github.com/air-verse/air@latest

COPY ./src .

RUN go mod download

EXPOSE 8080

CMD [ "air", "-c", ".air.toml"]