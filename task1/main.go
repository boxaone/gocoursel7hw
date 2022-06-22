package main

import "fmt"

func removeDup(arr *[]int, ind int) {

	for s := ind + 1; s < len(*arr); s++ {
		if (*arr)[ind] == (*arr)[s] {
			*arr = append((*arr)[:s], (*arr)[s+1:]...)
		}
	}
}

func main() {
	arr := []int{4, 1, 4, 5, -4, 6, 3, 8, 8}
	var result []int
	result = make([]int, len(arr))

	copy(result, arr)
	fmt.Println("Arr:    ", arr)

	for i := 0; i < len(result); i++ {
		removeDup(&result, i)
	}

	fmt.Println("Result: ", result)
}
