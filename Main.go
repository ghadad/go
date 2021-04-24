package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ghadad/firstgo/creature"
)

func main() {
	log.Println(os.Getenv("APP_ID"))
	fmt.Println(os.Getenv("GOPATH"))
	fmt.Printf("%d", os.Getpid())
	fmt.Println(creature.Hello())
	fmt.Println(creature.Add(4, 5))
}
