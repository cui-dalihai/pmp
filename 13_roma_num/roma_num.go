package main

import "fmt"

func romanToInt(s string) int {
	cnMap := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var res int
	var exception bool
	for k := range s {

		if exception {
			exception = false
			continue
		}

		if len(s)-1 > k {
			if cnMap[s[k]] < cnMap[s[k+1]] {
				res += cnMap[s[k+1]] - cnMap[s[k]]
				exception = true
				continue
			}
		}
		res += cnMap[s[k]]
	}
	fmt.Println(res)
	return res
}

func main() {

	romanToInt("XXVII")
}
