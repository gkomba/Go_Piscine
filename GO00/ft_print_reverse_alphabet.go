package main

import "fmt"

func ft_print_reverse_alphabet() {
	var c rune

	c = 'z'
	for c >= 'a' {
		fmt.Printf("%c", c)
		c--
	}
	fmt.Println()
}

// func main() {
// 	ft_print_reverse_alphabet()
// }
