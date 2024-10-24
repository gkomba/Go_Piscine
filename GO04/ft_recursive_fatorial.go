package main

import "fmt"

func	ft_recursive_fatorial(n int) int {
	if (n == 0 || n == 1) {
		return (1)
	}
	if n < 0 {
		return (0)
	}
	return (n * ft_recursive_fatorial(n - 1))
}

// func	main() {
// 	fmt.Println(ft_recursive_fatorial(5))
// }