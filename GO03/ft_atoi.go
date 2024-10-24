package main

import "fmt"

func ft_atoi(str string) int {
	var (
		sign int = 1
		res  int = 0
		base int = 10
		i    int = 0
	)

	for i < len(str) && str[i] == 32 {
		i++
	}

	for i < len(str) && (str[i] == '-' || str[i] == '+') {
		if str[i] == '-' {
			sign = -1
		}
		i++
	}

	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		res = res*base + int(str[i]-48)
		i++
	}
	return (res * sign)
}

func main() {
	fmt.Println(ft_atoi("1233"))
}
