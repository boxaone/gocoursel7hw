package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	input := "1 9 3 4 -5"

	var (
		result string
		min    int32
		max    int32
		num    int32
		numi   int
		err    error
	)

	for _, nums := range strings.Split(input, " ") {
		numi, err = strconv.Atoi(nums)

		if err == nil {
			num = (int32(numi))
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}

		}

	}

	switch {
	case min == max:
		result = fmt.Sprintf("%v", min)
	default:
		result = fmt.Sprintf("%v %v", max, min)

	}
	fmt.Printf("Input: %s\nOutput: %s", input, result)
}
