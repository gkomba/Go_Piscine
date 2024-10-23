package main

import "fmt"

func	ft_is_numeric(src string) int {
	for i := 0; i < len(src); i++ {
		if !(src[i] >= 48 && src[i] <= 57) {
			return 0
		}
	}
	return 1
}

// func	main() {
// 	fmt.Println(ft_is_numeric("1111"))
// }