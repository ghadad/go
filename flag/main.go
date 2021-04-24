package main

import (
	"flag"
	"fmt"
)

func main() {
	// variables declaration
	var uname string
	var pass string

	// flags declaration using flag package
	flag.StringVar(&uname, "u", "root", "Specify username. Default is root")
	flag.StringVar(&pass, "p", "password", "Specify pass. Default is password")

	flag.Parse() // after declaring flags we need to call it

	// check if cli params match
	if uname == "root" && pass == "password" {
		fmt.Printf("Logging in")
	} else {
		fmt.Printf("Invalid credentials!")
	}
}
