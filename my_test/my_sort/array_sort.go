package main

import (
	"fmt"
	"sort"
)

func main() {

	interval := [][]int{
		{2,3},
		{2,2},
		{3,3},
		{1,3},
		{5,7},
		{2,2},
		{4,6},
	}
	//interval 是二维数组 i 代表的是next j是前一个 TRUE的时候交换 false的时候不变动
	sort.Slice(interval, func(i, j int) bool {
		// 两行的第一位是相等的
		if  interval[i][0] ==  interval[j][0] {
			//第二个数值 大的放到前面
			// 按照第二列从大到小的顺序排序
			return interval[i][1] > interval[j][1] //false
		}else {
			// 按照第一列从小到大的顺序排序
			return interval[i][0] <  interval[j][0]
		}
	})
	fmt.Println(interval)

	//闭包的概念
}
