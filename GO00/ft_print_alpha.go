package main

import "fmt"

func ft_print_alpha() {

	var c rune
	c = 'a'
	for c <= 'z' {
		fmt.Printf("%c", c)
		c++
	}
	fmt.Println()
}

// func main() {
// 	ft_print_alpha()
// }
