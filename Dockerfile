# syntax=docker/dockerfile:1

FROM docker.io/golang:1.20.0-alpine3.17

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .
RUN go build -v -o samplegingo

EXPOSE 8080

CMD [ "/app/samplegingo" ]