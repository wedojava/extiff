package main

import (
	"fmt"
	"log"

	"github.com/wedojava/extiff"
)

var (
	cfg string = "./config.txt"
	src string = "./"
)

func init() {
	fmt.Println("[!] Input your config file path (`./config.txt` is default while you input nothing):")
	fmt.Print("[*] ->")
	fmt.Scanln(&cfg)
	fmt.Println("[!] Input path of tifs located (`./` is default while you input nothing):")
	fmt.Print("[*] ->")
	fmt.Scanln(&src)
	fmt.Printf("[config]: %s\n[tif at]: %s\n", cfg, src)
}

func main() {
	log.Println("[!] Many years later ...")
	_, err := extiff.Handle(cfg, src)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[!] Done.")
	// Remain for test
	// for _, t := range ts {
	//         fmt.Println(t)
	// }
}
