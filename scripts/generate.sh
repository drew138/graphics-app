# protoc -I src/ --go_out=src/ src/proto/messages.proto
protoc --go_out=src/ --go-grpc_out=src/ proto/messages.proto
