package main

import (
	"fmt"
	"strconv"
)

func main() {
	var array []uint = []uint{1, 1, 2, 4, 4, 5, 3, 3, 2, 6, 6}
	var array2 = []uint{1, 1, 2, 4, 5, 5, 3, 3, 2, 6, 6}
	array3 := []uint{1, 1, 5, 4, 4, 5, 3, 3, 2, 6, 6, 2}

	fmt.Println(queryOne(array))
	fmt.Println(queryOne(array2))
	fmt.Println(queryOne(array3))

	number := 12212
	var result bool = isPalindromeNumber(number)
	print(result)
}

func queryOne(array []uint) uint {
	counts := make(map[uint]int, 15)
	for _, k := range array {
		counts[k]++
	}
	for k, _ := range counts {
		if counts[k] == 1 {
			return k
		}
	}
	return 0
}

func isPalindromeNumber(number int) bool {
	if number < 0 {
		return false
	}
	str := strconv.Itoa(number)
	bytes := []byte(str)
	for i := 0; i < len(bytes)/2; i++ {
		if bytes[i] != bytes[len(bytes)-1-i] {
			return false
		}
	}
	return true
}
