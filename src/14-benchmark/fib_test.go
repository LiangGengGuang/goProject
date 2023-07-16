package main

import (
	"testing"
	"time"
)

// @Description
// @Author lianggengguang
// @Date 2023/7/15

/*
	再确认性能瓶颈点，度量某个函数或方法的性能
	benchmark性能测试	go test -bench .
	只运行以 Fib 结尾的 benchmark 用例	go test -bench='Fib$' .
	执行结果：
			goos: darwin
			goarch: amd64
			pkg: fib
			cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
			BenchmarkFib-12              332           3524772 ns/op
			PASS
			ok      fib     2.941s

	-12 即 GOMAXPROCS，默认等于 CPU 核数 可以通过 -cpu 参数改变 GOMAXPROCS，-cpu 支持传入一个列表作为参数
	go test -bench='Fib$' -cpu=2,4 .

	332           3524772 ns/op 表示用例执行了 332 次，每次花费约 0.003s

	benchmark 的默认时间是 1s，那么我们可以使用 -benchtime 指定为 5s
	go test -bench='Fib$' -benchtime=5s .

	调整具体的执行次数 -benchtime=50x
	go test -bench='Fib$' -benchtime=50x .

	-count 参数可以用来设置 benchmark 的轮数。例如，进行 3 轮 benchmark。
	go test -bench='Fib$' -benchtime=5s -count=3 .
*/
func BenchmarkFib(b *testing.B) {
	//属性 b.N 表示这个用例需要运行的次数
	//b.N 从 1 开始，如果该用例能够在 1s 内完成，b.N 的值便会增加，再次执行。b.N 的值大概以 1, 2, 3, 5, 10, 20, 30, 50, 100 这样的序列递增，越到后面，增加得越快
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}

/*
	-benchmem 参数可以度量内存分配的次数
	go test -bench='Generate' -benchmem
	执行结果：
			goos: darwin
			goarch: amd64
			pkg: example
			cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
			BenchmarkGenerateWithCap-12           76          15157307 ns/op         8003638 B/op          1 allocs/op
			BenchmarkGenerate-12                  57          20567318 ns/op        41678176 B/op         39 allocs/op
			PASS
			ok      example 4.650s

	Generate 分配的内存是 GenerateWithCap 的5倍，设置了切片容量，内存只分配一次，而不设置切片容量，内存分配了39次。

*/
func BenchmarkGenerateWithCap(b *testing.B) {
	time.Sleep(time.Second * 3) // 测试时有时会存在准备任务，模拟耗时准备任务
	b.ResetTimer()              // 重置定时器  	其他方法：b.StartTimer()、b.StopTimer()
	for n := 0; n < b.N; n++ {
		generateWithCap(1000000)
	}
}

func BenchmarkGenerate(b *testing.B) {

	for n := 0; n < b.N; n++ {
		generate(1000000)
	}
}
