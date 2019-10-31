package main

import (
	"fmt"
	"os"
)

func main() {
	var GOPATH, GOROOT string
	GOPATH = os.Getenv("GOPATH")
	GOROOT = os.Getenv("GOROOT")

	fmt.Println("GOPATH:", GOPATH)
	fmt.Println("GOROOT:", GOROOT)

}

