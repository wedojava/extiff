package extiff

import (
	"fmt"
	"log"
)

func ExampleReadArea() {
	cft := "./example/config.txt"
	as, err := ReadArea(cft)
	if err != nil {
		log.Fatalf("Read error occur: %v", err)
	}
	for _, a := range as {
		fmt.Println(a)
	}
	// Output:
	// {testArea1 {{123 131 719 819}}}
	// {testArea2 {{124 132 729 829}}}
	// {testArea3 {{125 133 739 839}}}
	// {testArea4 {{126 134 749 849}}}
	// {testArea5 {{127 135 759 859}}}
	// {testArea6 {{128 136 769 869}}}
	// {testArea7 {{129 137 779 879}}}
}
