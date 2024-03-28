FROM golang:latest

RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app

COPY ./ ./

ENTRYPOINT CompileDaemon -build="go build -o build/user-service" -command="./build/user-service" -build-dir="/app"