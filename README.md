# gRPC service that get Str and return LenOfStr
### commands for start (for macOS)

if you don't have installed protoc:
```bash
brew install protobuf
```

install plagins for Go
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest &&\
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

for plagins work
```bash
export PATH=$PATH:$(go env GOPATH)/bin  
```

```bash
protoc -I proto proto/LenOfStr/*.proto\
       --go_out=./gen/go/ --go_opt=paths=source_relative\
       --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
```