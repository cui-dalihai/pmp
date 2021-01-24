package main

import (
	"fmt"
	"sort"
)

// more ide test

// more test explain

type Node struct {
	val   int
	index int
}
type NodeList []*Node

// 实现排序接口
func (nl NodeList) Len() int           { return len(nl) }
func (nl NodeList) Less(i, j int) bool { return nl[i].val < nl[j].val }
func (nl NodeList) Swap(i, j int)      { nl[i], nl[j] = nl[j], nl[i] }

func twoSum(nums []int, target int) []int {

	// 每个slice元素转为节点类型
	// 节点存储原始val和index
	var nodes NodeList
	for index, val := range nums {
		nodes = append(nodes, &Node{index: index, val: val})
	}
	if nodes == nil {
		return nil
	}

	// 对节点按照val进行排序
	sort.Sort(nodes)

	start := 0
	end := len(nodes) - 1

	// 使用双游标查找目标
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
	return nil
}

func main() {
	al := []int{9, 3, 6, 1, 4, 12, 23, 7, 11}
	res := twoSum(al, 9)
	fmt.Println(res)
}
