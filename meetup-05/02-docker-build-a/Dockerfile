# Start from a golang image
FROM golang as binaryBuilder

# set working directory, this is the path you are inserted into the container at.
WORKDIR /go/src/helloworld

# copy in the source file
COPY $PWD/main.go .

# compile the binary
RUN GOOS=linux GOARCH=386 go build -v main.go

# Expose the port the container will run on. This should match the port in the binary.
EXPOSE 6060

# command to run on container start
CMD ["./main"]
