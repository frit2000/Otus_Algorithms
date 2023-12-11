package main

import (
	"fmt"
)

func buildArray(baseArray map[int]int) map[int]int {
	newLen := len(baseArray) + 9
	newArr := make(map[int]int, 0)
	for i := 0; i < newLen; i++ {
		val := 0
		for j := 0; j < 10; j++ {
			_, ok := baseArray[i-j]
			if ok {
				val += baseArray[i-j]
			}
		}
		newArr[i] = val
	}
	//	fmt.Println("newArr=", newArr)
	return newArr
}
func luckyTickets(n int) int {
	baseArr := make(map[int]int, 0)
	result := 0
	for i := 0; i < 10; i++ {
		baseArr[i] = 1
	}
	for i := 0; i < n-1; i++ {
		baseArr = buildArray(baseArr)
	}
	for _, a := range baseArr {
		result = result + a*a
	}
	return result

}

func main() {

	for n := 1; n <= 10; n++ {
		fmt.Println("Tickets of", n, "=", luckyTickets(n))
	}

}
