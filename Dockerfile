# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.17-alpine as build

WORKDIR /auth-service

COPY . /auth-service
RUN go mod download

WORKDIR /auth-service/app

RUN --mount=type=cache,target=/root/.cache CGO_ENABLED=0 go build -v -ldflags '-w -s' -o auth-service
RUN chmod 777 auth-service

# Run stage

FROM alpine:latest

WORKDIR /app
COPY --from=build /auth-service/app/auth-service .
EXPOSE 8080

CMD [ "/app/auth-service" ]