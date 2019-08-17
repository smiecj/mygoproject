package main

import "fmt"

func addOne(num *int) {
	*num = *num + 1
}

func main() {
	num := 1
	addOne(&num)
	fmt.Printf("1+1=%d\n", num)
}