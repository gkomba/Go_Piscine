package main

import "fmt"

func	ft_putstr(str *string) {
	fmt.Println(*str);
}

func	main() {
	var hello string;

	hello = "clear";
	ft_putstr(&hello);
}