FROM golang:alpine as builder

WORKDIR /app

RUN apk --no-cache add git
RUN apk --no-cache add build-base

COPY *go ./
COPY go.* ./

# If we don't specify the output filename with -o <filename>, go build will
# default to naming the output file after the project's directory name, in this
# case "shippy".  This confused the hell out of me.  Then I ran:
# `go help build` and learned about this.  I'd rather be explicit.  I hate 
# unspoken conventions.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consignment-client .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY consignment.json /app/consignment.json
# Note that the filename here, consignment-client, must match the name of the output
# file created in the above go build step.
COPY --from=builder /app/consignment-client .

ENTRYPOINT ["/app/consignment-client"]

# The name of the file to execute, must match the filename we built in the go build
# step above (this file has been copied into this container in the above COPY step)
CMD ["./consignment-client"]
