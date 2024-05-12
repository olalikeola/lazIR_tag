# lazIR_tag
cs 244b final project - distributed laser tag 

## Compile gRPC proto for Go
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"
```
```bash
cd src/proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./*.proto
```

## Compile gRPC proto in C/C++
Refer to `CmakeLists.txt`.