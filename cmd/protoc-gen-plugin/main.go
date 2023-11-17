// protoc-gen-plugin is a plugin for the Google protocol buffer compiler to
// generate Go code. Install it by building this program and making it
// accessible within your PATH with the name:
//
//	protoc-gen-go-grpc
//
// The 'go-grpc' suffix becomes part of the argument for the protocol compiler,
// such that it can be invoked as:
//
//	protoc --go-grpc_out=. path/to/file.proto
//
// This generates Go service definitions for the protocol buffer defined by
// file.proto.  With that input, the output will be written to:
//
//	path/to/file_grpc.pb.go
package main

import "fmt"

func main() {
	// Return a greeting that embeds the name in a message.
	fmt.Println("Hi, welcome!")
}
