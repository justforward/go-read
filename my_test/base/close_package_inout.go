package main

import (
	"fmt"
	"sync/atomic"
)

// 使用结构体统计所有应用的进出口流量，同时为了避免多个应用同时对结构体进行数据修改，必须保证原子操作。
// 保证原子的操作

// 使用全局变量实例化结构体
var bps = NewBps()

//针对不同的流量进出
type Bps struct {
	InByteAddF    func(uint64)  // 辅助字段，用来对入口流量进行操作
	InByteCountF  func() uint64 // 函数计算当前入口流量
	OutByteAddF   func(uint64)
	OutByteCountF func() uint64
}

// 工厂函数
func NewBps() *Bps {
	b := &Bps{}
	// 传入的流量
	b.InByteAddF, b.InByteCountF = addFunc()
	// 返回两个函数 添加的
	b.OutByteAddF, b.OutByteCountF = addFunc()

	return b
}

// 函数，返回两个函数，函数引用了外部变量，属于闭包
// 使用atomic包保证原子操作
func addFunc() (func(n uint64), func() uint64) {
	var count uint64 = 0
	// 返回了两个函数，函数引用了外部变量，属于闭包
	return func(n uint64) {
			// 添加 原子操作添加 n
			atomic.AddUint64(&count, n)
		}, func() uint64 {
			// 汇总
			c := count
			return c
		}
}

// 计算各个应用流量加入到总流量
func (b *Bps) Add(bytes uint64, in bool) {
	//进入的流量
	if in {
		// 调用传入的方法
		b.InByteAddF(bytes)
	} else {
		//出去的流量
		b.OutByteAddF(bytes)
	}

}

func Traffic(access bool, bytes uint64) {
	in := !access

	bps.Add(bytes, in)
}

// SampleBitrate,calculate all of network in and out traffic.
// unit is Bit
func SampleBitrate() (int64, int64) {
	return int64(bps.InByteCountF()), int64(bps.OutByteCountF())
}


func main() {
	type sub struct {
		//
		access bool
		data   uint64
	}

	testData := make([]sub, 6)
	// 代表不同的应用，收集进出口流量
	testData = append(testData, sub{true, 10})
	testData = append(testData, sub{false, 10})
	testData = append(testData, sub{false, 100})
	testData = append(testData, sub{false, 1000})
	testData = append(testData, sub{true, 1000})
	testData = append(testData, sub{true, 100})

	// 对 传入流量的原子操作
	for _, s := range testData {
		Traffic(s.access, s.data)
	}

	in, out := SampleBitrate()
	fmt.Printf("in:%d\tout:%d\n", in, out)

}
