package main

import	"fmt"

func	ft_ultimate_div_mod(a *int, b *int) {
	var	div int;
	var	mod	int;

	div = *a / *b;
	mod = *a % *b;
	*a = div;
	*b = mod;
}

// func	main() {
// 	a := 4;
// 	b := 2;
// 	fmt.Println(a);
// 	fmt.Println(b);
// 	ft_ultimate_div_mod(&a, &b)
// 	fmt.Println(a);
// 	fmt.Println(b);
// }