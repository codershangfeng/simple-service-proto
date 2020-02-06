# Proto file name 
PROTO?=""

.PHONY: generate
generate:
	protoc --go_out=plugins=grpc:.  $(PROTO)