package main

import "fmt"

func Repeat(ch string) string {
	const repeat = 5
	var res string
	i := 0
	for i < repeat {
		res += ch
		i++
	}
	return res
}

func main() {
	fmt.Println("hello")
}
