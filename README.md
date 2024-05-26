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

## Cross-compile for Raspberry Pi Zero
```bash
cd src/raft_grpc
env GOOS=linux GOARCH=arm64 GOARM=6 go build -o raft_grpc
```

## Start a raft cluster
1. Referring to config templates in `config`, prepare a config for each node.
2. Cross-compile `raft_grpc` for Raspberry Pi Zero(referring to [section](#cross-compile-for-raspberry-pi-zero))
3. SCP `raft_grpc` and config files to Raspberry Pi Zero
4. On Raspberry Pi Zero, run `raft_grpc --config path_to_config`
