## Installation
Make sure if you have protobuf installed in your system if not, use command below.
```bash
brew install protobuf
```

Make sure if you have protoc-gen-go installed in your system if not, use command below.
```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

Make sure if you have protoc-gen-grpc-gateway installed in your system if not, use command below.
```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

Make sure if you have protoc-gen-swagger installed in your system if not, use command below.
```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```
this will generate a ```swagger.json``` file

## Compile Proto files
before your run server or client, you should compliled .proto files using this command.
```bash
make
```

## Run Server
gRPC server will run at port ```:7777``` and REST server will run at port ```:7778```
```bash
go run server.go
```
### Run Client
```bash
go run client.go
```

## Test Rest API
```bash
curl -H "login:john" -H "password:doe" -X POST -d '{"to":"foo@foo@foo", "subject":"Foo", "message":"Foo"}' 'http://localhost:7778/1/email'
```