package leetcode

import (
	"fmt"
)

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4} // 声明了一个int类型的数组
	max := maxSubArray(arr)
	fmt.Printf("max = %d \t", max)
}

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	res := nums[0]
	for i := 1; i < l; i++ {
		nums[i] += max(nums[i-1], 0)
		res = max(nums[i], res)
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b

}

// 数组知识点补充：
func newArray() {
	//创建数组(声明长度)
	var array1 = [5]int{1, 2, 3}
	fmt.Printf("array1--- type:%T \n", array1)
	rangeIntPrint(array1[:])

	//创建数组(不声明长度)
	var array2 = [...]int{6, 7, 8}
	fmt.Printf("array2--- type:%T \n", array2)
	rangeIntPrint(array2[:])

	//创建数组切片
	var array3 = []int{9, 10, 11, 12}
	fmt.Printf("array3--- type:%T \n", array3)
	rangeIntPrint(array3)

	//创建数组(声明长度)，并仅初始化其中的部分元素
	var array4 = [5]string{3: "Chris", 4: "Ron"}
	fmt.Printf("array4--- type:%T \n", array4)
	rangeObjPrint(array4[:])

	//创建数组(不声明长度)，并仅初始化其中的部分元素，数组的长度将根据初始化的元素确定
	var array5 = [...]string{3: "Tom", 2: "Alice"}
	fmt.Printf("array5--- type:%T \n", array5)
	rangeObjPrint(array5[:])

	//创建数组切片，并仅初始化其中的部分元素，数组切片的len将根据初始化的元素确定
	var array6 = []string{4: "Smith", 2: "Alice"}
	fmt.Printf("array6--- type:%T \n", array6)
	rangeObjPrint(array6)
}

//输出整型数组切片
func rangeIntPrint(array []int) {
	for i, v := range array {
		fmt.Printf("index:%d  value:%d\n", i, v)
	}
}

//输出字符串数组切片
func rangeObjPrint(array []string) {
	for i, v := range array {
		fmt.Printf("index:%d  value:%s\n", i, v)
	}
}
