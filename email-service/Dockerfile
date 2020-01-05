# vessel-service/Dockerfile
FROM golang:alpine as builder

WORKDIR /app

RUN apk --no-cache add git
RUN apk --no-cache add build-base

COPY *go ./
COPY go.* ./
COPY proto ./proto/

RUN CGO_ENABLED=0 GOOS=linux go build -o vessel-service -a -installsuffix cgo main.go repository.go handler.go datastore.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/vessel-service .

CMD ["./vessel-service"]
