package main

import (
	"fmt"
)

func	ft_strcpy(src string, dest []byte) {
	for i, char := range src {
		dest[i] = byte(char)
	}
	dest[len(src)] = 0
}

// func main() {
// 	src := "Hello"
// 	dest := make([]byte, len(src)+1)

// 	ft_strcpy(src, dest)
// 	fmt.Println("Source:", src)
// 	fmt.Println("Destination:", string(dest))
// }