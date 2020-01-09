# vessel-service/Dockerfile
FROM golang:alpine as builder

WORKDIR /app

RUN apk --no-cache add git
RUN apk --no-cache add build-base

COPY *go ./
COPY go.* ./

RUN CGO_ENABLED=0 GOOS=linux go build -o shippy-email-service -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/shippy-email-service .

CMD ["./shippy-email-service"]
