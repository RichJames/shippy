# vessel-service/Dockerfile
FROM golang:alpine as builder

WORKDIR /go/src/github.com/RichJames/shippy/vessel-service

RUN apk --no-cache add git
RUN apk --no-cache add build-base

RUN mkdir -p /go/src/vessel-service && cd /go/src/vessel-service
RUN mkdir proto

COPY *go ./
COPY proto ./proto/

RUN go mod init
RUN go build ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o vessel-service -a -installsuffix cgo main.go repository.go handler.go datastore.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/RichJames/shippy/vessel-service .

CMD ["./vessel-service"]
