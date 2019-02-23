protobuild:
	protoc -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--proto_path=api --go_out=plugins=grpc:api api/*.proto \
	--grpc-gateway_out=logtostderr=true:api \
	--swagger_out=logtostderr=true:api