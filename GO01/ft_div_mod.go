package main

import "fmt"

func	ft_div_mod(a int , b int, div *int, mod *int) {
	*div = a / b; 
	*mod = a % b;
}

func	main() {
	a := 4;
	b := 2;
	div := 0;
	mod := 0;

	ft_div_mod(a, b, &div, &mod);
	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(div);
	fmt.Println(mod);
}