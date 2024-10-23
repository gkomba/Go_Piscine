package main

import "fmt"

func	ft_print_numbers() {
	var nbr int;

	nbr = 0;
	for nbr <= 9 {
		fmt.Printf("%d", nbr);
		nbr++
	}
	fmt.Println();
}

// func	main() {
// 	ft_print_numbers();
// }