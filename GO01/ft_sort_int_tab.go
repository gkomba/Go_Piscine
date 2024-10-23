package main

import "fmt"

func	ft_sort_int_tab(arr []int) {
	size := len(arr)
	for i := 0; i < size; i++ {
		if (arr[i] > arr[i + 1]) {
			tmp := arr[i]
			arr[i] = arr[i + 1]
			arr[i + 1] = tmp
			i = 0;
		} else {
		i++; 
	}
	}
}

// func	main() {
// 	arr := []int{2,1,3}
// 	fmt.Println(arr)
// 	ft_sort_int_tab(arr)
// 	fmt.Println(arr)
// }