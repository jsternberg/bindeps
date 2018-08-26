package main

//go:generate bin/protoc --plugin bin/protoc-gen-gogofaster --gogofaster_out=plugins=grpc:. ./hello.proto

import "fmt"

func main() {
	fmt.Println("vim-go")
}
