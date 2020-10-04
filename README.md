# Golang gRPC Chat

## Development
### Dependencies
Make sure to run `go mod vendor` to ensure all dependencies are within the projects vendor folder

### Protobuf
Make sure you have the following binaries installed
- `protoc`
    - macos: run `brew install protobuf`, make sure to install version 3.*
- `protoc-gen-go`
    - comes for free with the dependency `github.com/golang/protobuf`
    - make sure GOPATH binaries are part of your PATH (`export PATH="$PATH:$(go env GOPATH)/bin"`)

To generate the protobuf code run from the project root  
`$ protoc -I pb/ pb/pb.proto --go_out=plugins=grpc:./pb --go_opt=paths=source_relative`

### Configuration
For now all configuration is hardcoded in `config/constants.go`

### Running the applications
gRPC server that keeps history per chatroom  
`$ go run cmd/server/main.go`

Web(Socket) server that hosts the frontend and does some bi-directional message streaming with the gRPC server  
`$ go run cmd/chatroom/main.go`  

I wrote a simulation to get a grasp of how the default case of a select statement works in go.
I kept it here because it might help someone understand it better too.  
`$ go run cmd/simulate_disconnected_client/main.go`

## Packages / Folders
| Folder    | Contents |
|:----------|:---------|
| chat      | WebSocket client/server code and gRPC implementations |
| cmd       | Application entrypoints                               |
| config    | Application configuration                             |
| pb        | Protobuf files and generated proto code               |
| static    | Front-end code of chatroom app                        |
