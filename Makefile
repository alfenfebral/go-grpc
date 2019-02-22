protobuild:
	protoc --proto_path=api --go_out=plugins=grpc:api api/*.proto