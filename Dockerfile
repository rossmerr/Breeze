FROM golang:1.12-alpine3.9 AS builder
RUN apk add git

WORKDIR /go/src/github.com/rossmerr/breeze/cmd/server
COPY ./cmd/server/ . 
COPY ./go.mod . 
RUN GO111MODULE=on go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -a -o server .

FROM golang:1.12-alpine3.9
RUN apk --no-cache add ca-certificates
RUN mkdir /www
WORKDIR /root/
COPY --from=builder /go/src/github.com/rossmerr/breeze/cmd/server/server .
COPY --from=builder /go/src/github.com/rossmerr/breeze/cmd/server/index.html /www
COPY --from=builder /usr/local/go/misc/wasm/wasm_exec.js /www

VOLUME [ "/src" ]
EXPOSE 8080
CMD ["./server"]  
