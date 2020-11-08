package main

import (
	"fmt"
	"sort"
)

type node struct {
	val    int
	symbol string
}

var cnMap = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

type units []*node

func (us units) Len() int           { return len(us) }
func (us units) Less(i, j int) bool { return us[i].val > us[j].val }
func (us units) Swap(i, j int)      { us[i], us[j] = us[j], us[i] }

func sortedNodes(cnMap *map[int]string) units {
	nodes := units{}
	for k, v := range *cnMap {
		nodes = append(nodes, &node{val: k, symbol: v})
	}

	sort.Sort(nodes)
	return nodes
}

func intToRoman(num int) string {

	res := ""

	nodes := sortedNodes(&cnMap)
	for i := 0; i < len(nodes); i++ {

		if num < nodes[i].val {
			continue
		}

		num -= nodes[i].val
		res += nodes[i].symbol
		i -= 1
	}

	return res
}

func main() {

	//repeat := false
	//
	//al := []int{1,2,3,4,5,6,7}
	//for k, v := range al{
	//	fmt.Println(k, v)
	//	if k == 2 && !repeat {
	//		k = 1
	//		repeat = true
	//	}
	//}
	//
	//for i:=0; i<len(al); i++ {
	//	fmt.Println(i, al[i])
	//	if i==2 && !repeat {
	//		i -= 1
	//		repeat = true
	//	}
	//}

	fmt.Println(intToRoman(53))
}
