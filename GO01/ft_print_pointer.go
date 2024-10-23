package main

import "fmt"

func ft_print_pointer() {
	var a int
	var p *int

	a = 42
	p = &a

	fmt.Printf("var -> %d\n", a)
	fmt.Printf("pointer -> %d\n", *p)
}

// func main() {
// 	ft_print_pointer()
// }
