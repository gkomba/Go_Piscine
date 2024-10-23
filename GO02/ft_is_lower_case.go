package main

import "fmt"

func	ft_str_is_lower_case(src string) int {
	for i := 0; i < len(src); i++ {
		if !(src[i] >= 'a' && src[i] <= 'z') {
			return 0
	}
}
return 1
}

// func	main() {
// 	fmt.Println(ft_str_is_lower_case("qA"))
// }