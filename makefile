PHONY: generate
generate:
	protoc --go_out=./ --go-grpc_out=./ ./pkg/grpc_proto/profile.proto