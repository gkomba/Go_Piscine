package main

import "fmt"

func	ft_str_lower_case(str string) string {
	tmp_str := []byte(str)
	for i := 0; i < len(tmp_str); i++ {
		if (tmp_str[i] >= 'A' && tmp_str[i] <= 'Z') {
			tmp_str[i] += 32;
		}
	}
	return (string(tmp_str))
}

// func main() {
// 	fmt.Println(ft_str_lower_case("AAA"))
// }