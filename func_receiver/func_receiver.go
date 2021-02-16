package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// 函数和方法来自于不同的命名空间, 所以这两个命名不会冲突
// 但是如果为Point定义一个X方法, 就会冲突, 因为方法和字段来自于一个同一个命名空间
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 除了指针类型和接口类型的命名类型不允许声明方法, 其他的都可以
type P *int

// func (p *P) f() {}  // Invalid receiver type 'P' ('P' is a pointer type)

type Path []Point

// 不同类型有着各自独立的命名空间, 所以Path的Distance方法和Point的Distance方法不冲突
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// 由于调用一个函数时会复制每一个参数的值给这个函数
// 所以如果这个函数需要更新这些参数, 或者这些参数的值太大我们不想直接复制,
// 那就可以传递参数的指针给函数, 这同样适用于方法
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type Test struct {
	buf  []int
	name string
}

func (t *Test) update() {
	t.buf = []int{10, 10}
	t.name = "test"
}

type Test1 struct {
	buf  []int
	name string
}

func (t Test1) update() {
	t.buf = []int{10, 10}
	t.name = "test1"
}

func main() {
	aa := Test{[]int{0, 0}, "he"}
	fmt.Println(aa)
	bb := aa
	fmt.Printf("%p\n", &(aa.buf[0]))
	fmt.Printf("%p\n", &(bb.buf[0]))
	fmt.Println(bb)
	aa.update()
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println("-----")
	cc := Test1{[]int{1, 1}, "ha"}
	fmt.Println(cc)
	dd := cc
	fmt.Printf("%p\n", &(cc.buf))
	fmt.Printf("%p\n", &(dd.buf))
	fmt.Println(dd)
	cc.update()
	fmt.Println(cc)
	fmt.Println(dd)

	a := &Point{1, 2}
	a.ScaleBy(2)

	// ScaleBy方法接收者是个指针类型, 应该向上面那样使用指针类型作为接受者
	// 下面这种直接使用Point类型的变量作为接收者, 编译器会对变量进行隐式取址变换&b
	b := Point{1, 2}
	b.ScaleBy(2)    // 直接使用类型调用使用指针作为接收者的方法
	(&b).ScaleBy(2) // 显式取指针调用指针作为接收者的方法
	fmt.Println(a, b)

	// 对称的, Distance使用Point类型作为接受者
	// 但是像下面这种使用指针也是可以直接调用的, 编译器会隐式对指针进行取值变换*c
	c := &Point{2, 5}
	fmt.Println(c.Distance(b))    // 直接使用指针调用类型变量作为接收者的方法
	fmt.Println((*c).Distance(b)) // 显式取类型的变量调用类型变量作为接收者的方法

	// 综上, 无论方法使用指针还是类型作为接收者, 指针或者类型的变量都可以直接调用方法
	// 而如果要更新变量，则只能用指针作为接收者
}
