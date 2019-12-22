FROM golang:1.13.0 as builder

WORKDIR /go/src/github.com/RichJames/shippy/consignment-cli

COPY . .

#RUN apt-get update
#RUN apt-get install go-dep
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY consignment.json /app/consignment.json
COPY --from=builder /go/src/github.com/RichJames/shippy/consignment-cli/consignment-cli /app/consignment-cli

ENTRYPOINT ["/consignment-cli", "consignment.json"]
CMD ["./consignment-cli"]
