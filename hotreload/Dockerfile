FROM golang:1.12-alpine3.9

RUN apk add git
RUN go get github.com/githubnemo/CompileDaemon


VOLUME [ "/src" ]
WORKDIR /src

VOLUME [ "/public" ]

ENTRYPOINT CGO_ENABLED=0 GOOS=js GOARCH=wasm CompileDaemon -log-prefix=false -verbose=true -build="go build -a -o /public/breeze.wasm ." -command="cp $(go env GOROOT)/misc/wasm/wasm_exec.js /public/"
