## install library

```python
go get google.golang.org/grpc
```

## generate Protocol Buffers

```python
cd golang-grpc
protoc -I ./ proto/add.proto --go_out=plugins=grpc:.
```

## Run Server

```python
cd golang-grpc/server
go run main.go
```

## Run Client

```python
cd golang-grpc/client
go run main.go
```
