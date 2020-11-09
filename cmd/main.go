package main

import (
	"log"

	"github.com/wedojava/extiff"
)

func main() {
	_, err := extiff.Handle("./example/config.txt", "./example/test/test.tif")
	if err != nil {
		log.Fatal(err)
	}
	// Remain for test
	// for _, t := range ts {
	//         fmt.Println(t)
	// }
}
