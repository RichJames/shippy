# consignment-service/Dockerfile

# We use the official golang image, which contains all the
# correct build tools and libraries.  Notice `as builder`,
# this gives this container a name that we can reference later on.

FROM golang:alpine as builder

WORKDIR /app

RUN apk --no-cache add git
RUN apk --no-cache add build-base

# Copy the current code into our workdir
COPY *go ./
COPY go.* ./
COPY proto ./proto/

# Build the binary, with a few flags which will allow
# us to run this binary in Alpine.
RUN CGO_ENABLED=0 GOOS=linux go build -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go

# Here we're using a second FROM statement, which is strange,
# but this tells Docker to start a new build process with this
# image.
FROM alpine:latest

# Security related package, good to have.
RUN apk --no-cache add ca-certificates

# Same as before, create a directory for our app.
WORKDIR /app

# Here, instead of copying the binary from our host machine,
# we pull the binary from the container named `builder`, within
# this build context.  This reaches into our previous image, finds
# the binary we built, and pulls it into this container. Amazing!
#COPY --from=builder /app/shippy-service-consignment/consignment-service .
COPY --from=builder /app/consignment-service .

# Run the binary as per usual! This time with a binary build in a 
# separate container, with all of the correct dependencies and
# run time libraries.
CMD ["./consignment-service"]
