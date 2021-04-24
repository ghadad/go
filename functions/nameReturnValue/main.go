package main

import (
	"fmt"
)

func main() {
	fmt.Println(divby10(100), divby20(100))
}

func divby10(num int) (res int) {
	res = num / 10
	return res
}

func divby20(num int) (res int) {
	res = num / 20
	return

}
