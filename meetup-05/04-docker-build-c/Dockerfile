# Start from a golang image
FROM golang as buildBinary

# set working directory
WORKDIR /go/src/helloworld

# copy in source
COPY "$PWD"/main.go .

# compile binary
RUN GOOS=linux GOARCH=386 go build -v -o main main.go

# Starting from an scratch container
FROM scratch

# copy the binary built in the golang container into the scratch container
COPY --from=buildBinary /go/src/helloworld/main .

# Expose the port the container will run on. This should match the port in the binary.
EXPOSE 6060

# command to run on container start
CMD ["./main"]
