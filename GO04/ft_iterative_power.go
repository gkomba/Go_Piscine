package main

import "fmt"

func	ft_iterative_power(nbr int, pow int) int {
	res := 0

	res = nbr
	for pow > 1{
		res *=nbr
		pow--
	}
	if (pow == 0) {
		return 1
	}
	if pow < 0 {
		return 0
	}
	return (res)
}

// func	main() {
// 	fmt.Println(ft_iterative_power(2, 2))
// }