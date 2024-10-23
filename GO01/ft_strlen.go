package main

import "fmt"

func	ft_strlen(str string) int {
	var	i int;

	i = 0;
	for range str {
		i++;
	}
	return (i);
}

func	main() {
	fmt.Println(ft_strlen("hello"));
}