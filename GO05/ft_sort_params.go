package main

import (
	"fmt"
	"os"
)

func ft_strcmp(s1 string, s2 string) int {
	length := len(s1)
	if length > len(s2) {
		length = len(s2)
	}
	for i := 0; i < length; i++ {
		if s1[i] != s2[i] {
			return int(s1[i]) - int(s2[i])
		}
	}
	return len(s1) - len(s2)
}

func ft_swap_strings(s1 *string, s2 *string) {
	var tmp_str string

	tmp_str = *s1
	*s1 = *s2
	*s2 = tmp_str
}

func sort_params(args []string) {
	for i := 0; i < (len(args) - 1); i++ {
		if ft_strcmp(args[i], args[i+1]) > 0 {
			ft_swap_strings(&args[i], &args[i+1])
		}
	}
}

func main() {
	cliargs := os.Args[1:]

	for _, args := range cliargs {
		fmt.Println(args)
	}
	fmt.Println()
	fmt.Println()
	sort_params(cliargs)
	for _, args := range cliargs {
		fmt.Println(args)
	}
}
