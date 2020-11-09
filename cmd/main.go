package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/wedojava/extiff"
)

var (
	cfg string = "./config.txt"
	src string = "./"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("[!] Input your config file path: (`./config.txt` is default while you input nothing)")
	for input.Scan() {
		if line := input.Text(); line != "" {
			cfg = input.Text()
		}
		break
	}
	fmt.Println("[!] Input path of tifs located: (`./` is default while you input nothing)")
	for input.Scan() {
		if line := input.Text(); line != "" {
			src = input.Text()
		}
		break
	}
	_, err := extiff.Handle(cfg, src)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[+] Done.")
	// Remain for test
	// for _, t := range ts {
	//         fmt.Println(t)
	// }
}
