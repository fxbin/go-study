package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "the xxx object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
