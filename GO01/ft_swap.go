package main

import "fmt"

func ft_swap(a *int, b *int) {
	var tmp int

	tmp = *a
	*a = *b
	*b = tmp
}

// func main() {
// 	x := 42
// 	y := 21

// 	fmt.Println(x, y)
// 	ft_swap(&x, &y)
// 	fmt.Println(x, y)
// }
