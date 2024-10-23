package main

import "fmt"

func ft_strcmp(s1 string, s2 string) int {
	length := len(s1)
	if length > len(s2) {
		length = len(s2)
	}
	for i := 0; i < length; i++ {
		if s1[i] != s2[i] {
			return (int(s1[i]) - int(s2[i]))
		}
	}
	return (len(s1) - len(s2))
}

// func main() {
// 	fmt.Println(ft_strcmp("gildo", "komba"))
// }
