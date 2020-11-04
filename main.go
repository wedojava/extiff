package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	filename = "./example/test/L18/test.tif"
	if filename == "" {
		fmt.Printf("Usage: cootiff <filename>\n")
		return
	}
}
