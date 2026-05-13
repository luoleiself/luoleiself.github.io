package main

import "testing"

// window: D:\Go\bin\go.exe test -benchmem -run=^$ -bench ^(BenchmarkWithoutPool|BenchmarkWithPool)$ github.com/luoleiself/goNote/sync

// 测试的函数名			共执行的次数即 b.N 的值			平均每次操作花费的时间 ns/op 纳秒			平均每次操作的内存申请数 B/op Byte			每次操作申请了多少次内存 alloc/op

func BenchmarkWithoutPool(b *testing.B) {
	var s *Student
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100000; j++ {
			s = new(Student)
			s.Name = "Tom"
			s.Age = 19
		}
	}
	b.Logf("s is %v\n", s)
}

func BenchmarkWithPool(b *testing.B) {
	var s *Student
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100000; j++ {
			s = StudentPool.Get().(*Student)
			s.Name = "Jerry"
			s.Age = 18
			StudentPool.Put(s)
		}
	}
}
