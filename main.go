package main

import (
	"fmt"

	"github.com/aikaiqiang/Algorithms-LeetCode-Go/kaywall"
)

const Pi = 3.1415926
const MaxThread = 10
const PREFIX = "astaxie_"

const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a       = iota //a=0
	b       = "B"
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)

//返回 A+B 和 A*B
func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func main() {
	kaywall.Hello()

	var a = 65
	b := string(a)
	fmt.Println("b值=", b)

	var comp complex64 = 5 + 5i
	fmt.Printf("Value is: %v\n", comp)

	s := "hello"
	by := []byte(s) // 将字符串 s 转换为 []byte 类型
	by[0] = 'A'
	s1 := string(by) // 再转换回 string 类型
	fmt.Printf("%s\n", s1)

	s2 := `test
        new line
	haha`
	fmt.Printf("%s\n", s2)

	// const
	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)

	var arr [10]int                                 // 声明了一个int类型的数组
	arr[0] = 42                                     // 数组下标是从0开始的
	arr[1] = 13                                     // 赋值操作
	fmt.Printf("The first element is %d\n", arr[0]) // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9])  //返回未赋值的最后一个元素，默认返回0

	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C") // 删除key为C的元素

	/**
	make、new操作
	make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。

	内建函数new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：
	new返回指针。
	内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T。本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。例如，一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；在这些项目被初始化之前，slice为nil。对于slice、map和channel来说，make初始化了内部的数据结构，填充适当的值
	*/

	X := 1
	Y := 2
	xPLUSy, xTIMESy := SumAndProduct(X, Y)

	fmt.Printf("%d + %d = %d\n", X, Y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", Y, Y, xTIMESy)
}
