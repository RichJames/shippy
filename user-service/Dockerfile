FROM golang:alpine as builder

#WORKDIR /go/src/github.com/RichJames/shippy/user-service
WORKDIR /app

RUN apk --no-cache add git
RUN apk --no-cache add build-base

COPY *go ./

#RUN go mod init .
#RUN go build ./...

# If we don't specify the output filename with -o <filename>, go build will
# default to naming the output file after the project's directory name, in this
# case "shippy".  This confused the hell out of me.  Then I ran:
# `go help build` and learned about this.  I'd rather be explicit.  I hate 
# unspoken conventions.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-service .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

# Note that the filename here, consignment-client, must match the name of the output
# file created in the above go build step.
COPY --from=builder /go/src/github.com/RichJames/shippy/consignment-cli/user-service .

# The name of the file to execute, must match the filename we built in the go build
# step above (this file has been copied into this container in the above COPY step)
CMD ["./user-service"]
