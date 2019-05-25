FROM golang:1.12-alpine3.9 AS builder

WORKDIR /go/src/github.com/rossmerr/breeze/cmd/http
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -a -o http .

FROM alpine:3.9  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/rossmerr/breeze/cmd/http .
VOLUME [ "/public" ]
EXPOSE 3000
CMD ["./http", "--public", "/public"]  
