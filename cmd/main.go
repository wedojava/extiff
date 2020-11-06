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
		for _, a := range as {
			if a.Env.Intersects(t.Env) {
				fmt.Println("matched one")
			} else {
				fmt.Println("not matched")
			}
		}
	}
	fmt.Println(ts) // why?
}
