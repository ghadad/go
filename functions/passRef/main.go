package main

import (
	"fmt"
)

func main() {
	val := 12
	fmt.Printf("The value before function call is %d\n", val)
	changeValue(val) // does not change value
	fmt.Printf("The value after function call is %d\n", val)
	changeValueByRef(&val) //changes value
	fmt.Printf("The value after function call is %d\n", val)
}

func changeValue(num int) {
	num = 42
	fmt.Printf("in function num is %d\n", num)
}

func changeValueByRef(num *int) {
	*num = 42
}
