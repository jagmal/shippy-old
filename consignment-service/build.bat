protoc -I. --go_out=plugins=grpc:. proto/consignment/consignment.proto
set GOARCH=amd64
set GOOS=linux
go build
