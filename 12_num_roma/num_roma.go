package main

import "fmt"

func intToRoman(num int) string {

	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romas := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""

	for i := 0; i < 13; i++ {
		if num < nums[i] {
			continue
		}
		num -= nums[i]
		res += romas[i]
		i -= 1
	}
	return res
}

func main() {
	fmt.Println(intToRoman(671))
}
