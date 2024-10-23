package main

import (
	"fmt"
)

func ft_strncpy(src string, dest []byte, n int) {
	var i int

	for i = 0; i < n && i < len(src); i++ {
		dest[i] = byte(src[i])
	}
	if n > 0 {
		dest[i] = 0
	}
}

// func main() {
// 	src := "Hello"
// 	dest := make([]byte, len(src)+1)

// 	ft_strncpy(src, dest, 2)
// 	fmt.Println("Source:", src)
// 	fmt.Println("Destination:", string(dest))
// }