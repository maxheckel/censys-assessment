FROM golang:latest

ENV APP_NAME censys
ENV PORT 8080

COPY . /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go get ./...


RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
EXPOSE 8000
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o ./build/${APP_NAME} ./cmd/app/" -command="./build/${APP_NAME}"

EXPOSE ${PORT}
