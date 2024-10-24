package main

import "fmt"

func	ft_fibonacii (index int) int {
	if (index < 0) {
		return -1
	}
	if (index < 2) {
		return (index)
	}
	return (ft_fibonacii(index - 2) + ft_fibonacii((index - 1)))
}

// func	 main() {
// 	fmt.Println(ft_fibonacii(4))
// }