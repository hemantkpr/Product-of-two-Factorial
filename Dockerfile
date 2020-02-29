FROM golang:1.13.7 as builder
WORKDIR /go/src/github.com/hemantkpr/product-of-two-factorial
RUN go get -d -v github.com/julienschmidt/httprouter
COPY main.go  .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/hemantkpr/product-of-two-factorial/main .
CMD ["./main"]