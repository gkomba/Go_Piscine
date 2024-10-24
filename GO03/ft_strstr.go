package main

import "fmt"

func ft_strstr(str string, to_find string) string {
	if len(to_find) == 0 {
		return str
	}

	for i := 0; i <= len(str)-len(to_find); i++ {
		if str[i:i+len(to_find)] == to_find {
			return (to_find)
		}
	}
	return ""
}

// func main() {
// 	fmt.Println(ft_strstr("ola mea caro", "meu"))
// }
