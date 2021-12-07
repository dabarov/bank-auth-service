# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /auth-service

COPY . /auth-service
RUN go mod download

WORKDIR /auth-service/app

RUN go build -o auth-service
RUN chmod 777 auth-service

EXPOSE 8080

CMD [ "./auth-service" ]