package main

import (
	"fmt"
	"sort"
)

type Node struct {
	val   int
	index int
}
type NodeList []*Node

func (nl NodeList) Len() int           { return len(nl) }
func (nl NodeList) Less(i, j int) bool { return nl[i].val < nl[j].val }
func (nl NodeList) Swap(i, j int)      { nl[i], nl[j] = nl[j], nl[i] }

func twoSum(nums []int, target int) []int {

	var nodes NodeList
	for index, val := range nums {
		nodes = append(nodes, &Node{index: index, val: val})
	}
	if nodes == nil {
		return nil
	}
	sort.Sort(nodes)

	start := 0
	end := len(nodes) - 1

	for start < end {
		if nodes[start].val+nodes[end].val == target {
			return []int{nodes[start].index, nodes[end].index}
		} else if nodes[start].val+nodes[end].val < target {
			for nodes[start].val == nodes[start+1].val {
				start++
			}
			start++
		} else {
			for nodes[end].val == nodes[end-1].val {
				end--
			}
			end--
		}
	}
	return []int{}
}

func main() {
	al := []int{9, 3, 6, 1, 4, 12, 23, 7, 11}
	res := twoSum(al, 9)
	fmt.Println(res)
}
