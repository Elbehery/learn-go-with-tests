package main

import "fmt"

func Repeat(ch string, n int) string {
	var res string
	i := 0
	for i < n {
		res += ch
		i++
	}
	return res
}

func main() {
	fmt.Println("hello")
}
