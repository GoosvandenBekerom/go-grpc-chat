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
`$ protoc -I chat/ chat/chat.proto --go_out=plugins=grpc:./chat --go_opt=paths=source_relative`

### Configuration
For now all configuration is hardcoded in `config/constants.go`

### Running the applications
To start the webserver with websocket support (no gRPC yet): `$ go build cmd/chat/main.go` 
To start the server: `$ go build cmd/server/main.go`  
To start a client: `$ go build cmd/client/main.go`
To start the "client disconnects" simulation: `$ go build cmd/simulate_disconnected_client/main.go`
