package main

import "fmt"

func ft_str_capitalize(str string) string {
	tmp_str := []byte(str)
	for i := 0; i < len(tmp_str); i++ {
		if i == 0 && tmp_str[i] >= 'a' && tmp_str[i] <= 'z' {
			tmp_str[i] -= 32
		}
		if i > 0 && tmp_str[i] >= 'a' && tmp_str[i] <= 'z' {
			if !(tmp_str[i-1] >= 'a' && tmp_str[i-1] <= 'z') && !(tmp_str[i-1] >= 'A' && tmp_str[i-1] <= 'Z') && !(tmp_str[i-1] >= '0' && tmp_str[i-1] <= '9') {
				tmp_str[i] -= 32
			}
		}
	}
	return (string(tmp_str))
}

// func main() {
// 	fmt.Println(ft_str_capitalize("oi, tudo bem? 42palavras quarenta-e-duas; cinquenta+e+um"))
// }
