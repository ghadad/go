package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "n", "admin", "Specify name. Default is admin.")

	flag.Usage = func() {
		fmt.Printf("Usage of our Program: \n")
		fmt.Printf("./go-project -n username\n")
		// flag.PrintDefaults()  // prints default usage
	}
	flag.Parse()
}
