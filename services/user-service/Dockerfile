FROM golang:latest AS development

RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /services/user-service

COPY ./ ./

RUN go mod download

ENTRYPOINT CompileDaemon -build="go build -o build/user-service" -command="./build/user-service" -build-dir="/services/user-service"