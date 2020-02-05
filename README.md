# simple-service-proto
The shared protobuf repository for simple service, and should be used as a submodule of git.

## Prons & Cons

- Prons:
    - An independent shared library between server and client.
    - Client could implement the gRPC protocol without access to the server repository.
    - The unique source of truth for gRPC protocl (Each commit should consider about back-compatibility)

- Cons:
    - (TODO)

## Key Commands

- Generate Code
    - Golang: `make generate`