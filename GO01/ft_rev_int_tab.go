package main

import "fmt"

func	 ft_rev_int_tab(arr []int) {
	size := len(arr)
	for i := 0; i < size / 2; i++{
		tmp := arr[i]
		arr[i] = arr[size - i - 1]
		arr[size - i -1] = tmp
	}
}

// func	main() {
// 	arr := []int{1,2,3}
// 	fmt.Println(arr)
// 	ft_rev_int_tab(arr)
// 	fmt.Println(arr)
// }