package main

import (
	"context"

	"github.com/gravitational/teleport/api/client"

	"fmt"
)

func main() {
	fmt.Println("hello world!")

	_, err := client.New(context.TODO(), client.Config{})
	if err != nil {
		panic(err)
	}
}
