package main

import (
	"fmt"
)

func main() {
	var a = new(int)
	*a = 1

	var b *int = a
	*b = 2
	fmt.Println(*a)
	fmt.Println(*b)

}
