package main

import "fmt"

func	ft_iterative_fatorial(n int) int {
	var res uint

	res = 1
	for n >  0 {
		res *= uint(n)
		n--	
	}
	if n < 0 {
		return (0)
	}
	return int(res)
}

// func	main() {
// 	fmt.Println(ft_iterative_fatorial(5))
// }