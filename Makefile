.PHONY: generate
generate:
	protoc --go_out=plugins=grpc:. userservice.proto