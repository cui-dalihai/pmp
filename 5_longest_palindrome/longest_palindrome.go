package main

import "fmt"

func IsPalindrome(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}
	hf := l / 2
	for i := 0; i <= hf; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}

// LongestPalindrome1 遍历所有子串
func LongestPalindrome1(s string) string {

	var runner, walker, maxlen int
	var maxstr string

	length := len(s)
	if length == 0 {
		return ""
	}

	for walker != length+1 {
		runner = walker
		for runner != length+1 {
			if IsPalindrome(s[walker:runner]) {
				if maxlen < runner-walker {
					maxlen = runner - walker
					maxstr = s[walker:runner]
				}
			}
			runner += 1
		}
		walker += 1
	}
	return maxstr
}

// LongestPalindrome2 中心扩散
func LongestPalindrome2(s string) string {

	var (
		maxlen, scan, length, curlen int
		maxstr                       string
	)

	length = len(s)
	if length == 0 {
		return ""
	}

	for scan != length {
		l := scan
		r := scan

		// 向左查找相同 l到相同字符左边界, 不包含边界
		for l >= 0 && s[scan] == s[l] {
			if maxlen < scan-l+1 {
				maxlen = scan - l + 1
				maxstr = s[l : scan+1]
			}
			l--
		}

		// 向右查找相同 r到相同字符右边界, 不包含边界
		for r < length && s[scan] == s[r] {
			if maxlen < r-scan+1 {
				maxlen = r - scan + 1
				maxstr = s[scan : r+1]
			}
			r++
		}

		// 同时向左向右查找对称, 此时l和r是移动后的
		for l >= 0 && r < length && s[l] == s[r] {
			curlen += 2
			if maxlen < r-l+1 {
				maxlen = r - l + 1
				maxstr = s[l : r+1]
			}
			l--
			r++
		}
		scan += 1
	}
	return maxstr
}

func main() {
	fmt.Println(LongestPalindrome2("bacabab"))
}
