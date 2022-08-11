STATUS:= `git status -s`

go:
	protoc -I=$(GOPATH)/src:. --gogofaster_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,paths=source_relative:. proto/v1/dto/*.proto
	protoc -I=$(GOPATH)/src:. --gogofaster_out=plugins=grpc,paths=source_relative:. proto/v1/opt/*.proto
	protoc -I=$(GOPATH)/src:. --gogofaster_out=plugins=grpc,paths=source_relative:. proto/v1/app/*.proto
	go test -v