package main

func ft_sqrt(nbr int) int64 {
	var index int64
	var tmp int64

	tmp = int64(nbr)
	if tmp <= 0 {
		return 0
	}
	if tmp == 1 {
		return 1
	}
	index = 2
	if tmp >= 2 {
		for index*index <= tmp {
			if index*index == tmp {
				return index
			}
			index++
		}
	}
	return 0
}

// func	main() {
// 	fmt.Println(ft_sqrt(4))
// }
