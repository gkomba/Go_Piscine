package main

import "fmt"

func ft_print_reverse_number() {
	var nbr int

	nbr = 9
	for nbr >= 0 {
		fmt.Printf("%d", nbr)
		nbr--
	}
	fmt.Println()
}

// func main() {
// 	ft_print_reverse_number()
// }
