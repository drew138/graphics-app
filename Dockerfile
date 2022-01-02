FROM golang:1.17

WORKDIR $GOPATH/src/github.com/Drew138/graphics-app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .
# https://stackoverflow.com/questions/66371020/how-to-use-docker-to-generate-grpc-code-based-on-go-mod-versions
RUN apt-get update

RUN apt install -y protobuf-compiler

RUN protoc --version

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
# EXPOSE 8080

# Run the executable
CMD ["go", "run", "main.go"]
