package main

import (
	"context"
	"fmt"

	"github.com/gravitational/teleport/api/client"

	"github.com/alexbathome/bazel-build-protobuf/internal/utils"
	pb "github.com/alexbathome/bazel-build-protobuf/proto/helloworld"
)

func main() {
	fmt.Println("hello world!")

	x := pb.HelloRequest{}
	x.Name = "test"

	fmt.Printf("%s\n", x.Name)

	z := 10
	r, _ := utils.TakeFive(z)
	fmt.Printf("res %d\n", r)

	_, err := client.New(context.TODO(), client.Config{})
	if err != nil {
		panic(err)
	}
}
