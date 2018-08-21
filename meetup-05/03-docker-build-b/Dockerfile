# Start from a golang image
FROM golang as binaryBuilder

# set working directory
WORKDIR /go/src/helloworld

# copy in source
COPY $PWD/main.go .

# compile binary
RUN GOOS=linux GOARCH=386 go build -v main.go

# Starting from an Alpine container
FROM alpine

# copy the binary built in golang container into the alpine container
COPY --from=binaryBuilder /go/src/helloworld/main /opt/helloworld/

WORKDIR /opt/helloworld/

# Expose the port the container will run on. This should match the port in the binary.
EXPOSE 6060

# command to run on container start
CMD ["./main"]
