package main

func LongestSubString(s string) (maxlen int) {
	var runner, walker int
	length := len(s)
	appeared := map[byte]int{}

	for runner != length {
		if v, ok := appeared[s[runner]]; ok && walker <= v {
			walker = v + 1
		} else {
			if maxlen < runner-walker+1 {
				maxlen = runner - walker + 1
			}
		}

		appeared[s[runner]] = runner
		runner += 1
	}
	return
}
