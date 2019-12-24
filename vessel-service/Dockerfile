# vessel-service/Dockerfile
FROM golang:alpine as builder

#WORKDIR /app/shippy-service-vessel
WORKDIR /go/src/github.com/RichJames/shippy/vessel-service

COPY . .

RUN apk --no-cache add bash
#RUN go mod download
#RUN go get -u github.com/golang/dep/cmd/dep
#RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vessel-service .


FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN apk --no-cache add bash

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/RichJames/shippy/vessel-service .

CMD ["./vessel-service"]
#CMD ["bash"]
