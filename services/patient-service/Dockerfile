FROM golang:latest AS development

RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app

COPY ./ ./

RUN go mod download

ENTRYPOINT CompileDaemon -build="go build -o build/patient-service" -command="./build/patient-service" -build-dir="/app"