# vessel-service/Dockerfile
FROM golang:alpine as builder

WORKDIR /app

RUN apk --no-cache add git
RUN apk --no-cache add build-base

COPY *go ./
COPY go.* ./

RUN CGO_ENABLED=0 GOOS=linux go build -o email-service -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/email-service .

CMD ["./email-service"]
