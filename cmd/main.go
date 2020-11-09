package main

import (
	"fmt"
	"log"

	"github.com/wedojava/extiff"
)

func main() {
	as, err := extiff.ReadArea("./example/config.txt")
	ts, err := extiff.GetTifs("./example")
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range ts {
		t.Extract()
		t.SetArea(as)
	}
	// Remain for test
	for _, t := range ts {
		fmt.Println(t.Areas)
	}
}
