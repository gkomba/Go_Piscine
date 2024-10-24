package main

import "fmt"

// func	main() {
// 	fmt.Println(ft_recursive_power(2, 2))
// }

func	ft_recursive_power(nbr int, pow int) (res int)  {
	res = nbr
	if (pow > 0) {
		return (nbr * ft_recursive_power(nbr, pow - 1))
	}
	if pow == 0 {
		return 1
	}
	if (pow < 0) {
		return 0
	}
	return (res)
}