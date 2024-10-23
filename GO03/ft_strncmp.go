package main

import "fmt"

func ft_strncmp(s1 string, s2 string, n int) int {
	length := len(s1)
	if length > len(s2) {
		length = len(s2)
	}
	for i := 0; i < length && i < n; i++ {
		if s1[i] != s2[i] {
			return (int(s1[i]) - int(s2[i]))
		}
	}
	return 0
}

// func main() {
// 	fmt.Println(ft_strncmp("gildo", "gggkomba", 2))
// }
