package hello

//go:generate bin/protoc --plugin bin/protoc-gen-gogofaster --gogofaster_out=plugins=grpc:. ./hello.proto
