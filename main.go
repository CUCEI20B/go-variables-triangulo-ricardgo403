package main

import "fmt"

func main() {
	var base uint64
	var altura uint64
	fmt.Scan(&base)
	fmt.Scan(&altura)
	area := (base * altura) / 2
	fmt.Print(area)
}
