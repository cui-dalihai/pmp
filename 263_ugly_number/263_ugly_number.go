package main


func isUgly(num int) bool {

	if num == 1 {
		return true
	}
	if num == 0 {
		return false
	}

	fs := []int{2,3,5}

	for _, v := range fs {
		for num % v == 0 {
			num = num / v
			if num == 1 {
				return true
			}
		}
	}
	return false
}


func main() {

}

