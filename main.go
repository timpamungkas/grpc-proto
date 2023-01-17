package main

import (
	"fmt"

	"github.com/timpamungkas/course-grpc-proto/protogen/go/hello"
)

func main() {
	req := &hello.HelloRequest{
		Name: "Tim",
	}

	fmt.Println(req)
}
