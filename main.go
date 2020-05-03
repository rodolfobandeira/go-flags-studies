package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	username := flag.String("u", "", "-u <username>")

	flag.Parse()

	if *username == "" {
		fmt.Println("Error! Check --help for usage parameters")
		os.Exit(1)
	}

	fmt.Println(*username)
}
