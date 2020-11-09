package main

import (
	"fmt"
	"log"

	"github.com/wedojava/extiff"
)

func main() {
	ts, err := extiff.Handle("./example/config.txt", "./example/test/L18/test.tif")
	if err != nil {
		log.Fatal(err)
	}
	// Remain for test
	for _, t := range ts {
		fmt.Println(t)
	}
}
