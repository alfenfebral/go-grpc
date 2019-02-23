## Run Server
```bash
go run server.go
```
### Run Client
```bash
go run client.go
```

## Usage
```bash
make
```
to compile proto files to go files

## Test Rest API
```bash
curl -H "login:john" -H "password:doe" -X POST -d '{"to":"foo@foo@foo", "subject":"Foo", "message":"Foo"}' 'http://localhost:7778/1/email'
```