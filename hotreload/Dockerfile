FROM golang:1.12-alpine3.9

RUN apk add git
RUN go get github.com/githubnemo/CompileDaemon


VOLUME [ "/public" ]
WORKDIR /public

ENTRYPOINT CGO_ENABLED=0 GOOS=linux CompileDaemon -log-prefix=false -verbose=true -build="go build -a -o wasm ." -command="./wasm"


