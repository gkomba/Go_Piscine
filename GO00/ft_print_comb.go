package main

import "fmt"

func ft_print_comb() {
	var a int
	var b int
	var c int

	a = 0
	for a <= 7 {
		b = a + 1
		for b <= 8 {
			c = b + 1
			for c <= 9 {
				fmt.Printf("%d%d%d", a, b, c)
				fmt.Println()
				c++
			}
			b++
		}
		a++
	}
}

// func main() {
// 	ft_print_comb()
// }
