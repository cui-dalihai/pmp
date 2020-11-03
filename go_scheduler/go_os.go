package main

import (
	"fmt"
	"runtime"
)

func main() {

	// 物理核心数6，超线程技术，共计12
	fmt.Println(runtime.NumCPU())

}
