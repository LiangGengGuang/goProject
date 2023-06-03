
#go get -u google.golang.org/protobuf/cmd/protoc-gen-go
#go install google.golang.org/protobuf/cmd/protoc-gen-go
#go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
#go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

echo '生成 rpc server 代码'

OUT=../server/rpc
protoc \
--go_out=${OUT} \
--go_opt=paths=source_relative \
--go-grpc_out=${OUT} \
--go-grpc_opt=paths=source_relative \
server.proto

echo '生成 rpc client 代码'

OUT=../client/rpc
protoc \
--go_out=${OUT} \
--go_opt=paths=source_relative \
--go-grpc_out=${OUT} \
--go-grpc_opt=paths=source_relative \
server.proto