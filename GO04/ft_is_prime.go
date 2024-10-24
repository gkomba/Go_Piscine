package main

import "fmt"

func ft_is_prime(n int) int {
	i := 2
	if n <= 1 {
		return 0
	}
	for i <= (n / 2) {
		if (n % i) == 0 {
			return 0
		} else {
			i += 1
		}
	}
	return 1
}

// func main() {
// 	fmt.Println(ft_is_prime(2))
// }
