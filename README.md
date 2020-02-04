# simple-service-proto
The shared protobuf repository for simple service, and should be used as a submodule of git.

## Key Commands

- Generate Code
    - Golang: `protoc --go_out=plugins=grpc:. [Path of Proto file]`