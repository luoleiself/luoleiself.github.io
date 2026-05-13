package internaltesting

import (
	// "bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// 单元(功能)测试
func testGetArea(t *testing.T) {
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	}
}
func testGetArea2(t *testing.T) {
	area := GetArea(10, 20)
	if area != 200 {
		t.Error("测试失败")
	}
}
func testGetArea3(t *testing.T) {
	area := GetArea(60, 70)
	if area != 4200 {
		t.Error("测试失败")
	}
}
func testGetArea4(t *testing.T) {
	area := GetArea(80, 90)
	if area != 7200 {
		t.Error("测试失败")
	}
}

// 包级别测试固件
func pkgSetUp(name string) func() {
	fmt.Printf("package setUp execute for %s\n", name)
	return func() {
		fmt.Printf("package tearDown execute for %s\n", name)
	}
}

// go 1.4 添加 M 通过 TestMain 函数去执行测试
func TestMain(m *testing.M) {
	defer pkgSetUp("TestMain")()
	m.Run()
}

// 测试固件
func setUp(name string) func() {
	fmt.Printf("\tsetUp execute for %s\n", name)
	return func() {
		fmt.Printf("\ttearDown execute for %s\n", name)
	}
}

/*
go test -v .
package setUp execute for TestMain
=== RUN   TestFunc1
				setUp execute for TestFunc1
=== RUN   TestFunc1/testGetArea
=== RUN   TestFunc1/testGetArea2
				tearDown execute for TestFunc1
--- PASS: TestFunc1 (0.00s)
		--- PASS: TestFunc1/testGetArea (0.00s)
		--- PASS: TestFunc1/testGetArea2 (0.00s)
=== RUN   TestFunc2
				setUp execute for TestFunc2
=== RUN   TestFunc2/testGetArea3
=== RUN   TestFunc2/testGetArea4
				tearDown execute for TestFunc2
--- PASS: TestFunc2 (0.00s)
		--- PASS: TestFunc2/testGetArea3 (0.00s)
		--- PASS: TestFunc2/testGetArea4 (0.00s)
PASS
package tearDown execute for TestMain
ok      github.com/luoleiself/learning-go/testing   (cached)
*/

func TestFunc1(t *testing.T) {
	// defer setUp(t.Name())() // go 1.14 之前使用
	t.Cleanup(setUp(t.Name())) // go 1.14 添加
	t.Run("testGetArea", testGetArea)
	t.Run("testGetArea2", testGetArea2)
}

func TestFunc2(t *testing.T) {
	// defer setUp(t.Name())() // go 1.14 之前使用
	t.Cleanup(setUp(t.Name())) // go 1.14 添加
	t.Run("testGetArea3", testGetArea3)
	t.Run("testGetArea4", testGetArea4)
}

/*
go test -v .
=== RUN   TestTDTFunc3
	setUp execute for TestTDTFunc3
=== RUN   TestTDTFunc3/TDTGetArea1
=== RUN   TestTDTFunc3/TDTGetArea2
=== RUN   TestTDTFunc3/TDTGetArea3
=== RUN   TestTDTFunc3/TDTGetArea4
	tearDown execute for TestTDTFunc3
--- PASS: TestTDTFunc3 (0.00s)
	--- PASS: TestTDTFunc3/TDTGetArea1 (0.00s)
	--- PASS: TestTDTFunc3/TDTGetArea2 (0.00s)
	--- PASS: TestTDTFunc3/TDTGetArea3 (0.00s)
	--- PASS: TestTDTFunc3/TDTGetArea4 (0.00s)
PASS

go test -v .
=== RUN   TestTDTFunc3
        setUp execute for TestTDTFunc3
=== RUN   TestTDTFunc3/TDTGetArea1
=== RUN   TestTDTFunc3/TDTGetArea2
    calc_test.go:153: [TDTGetArea2], want 2428.0800, but result is 1214.0400
=== RUN   TestTDTFunc3/TDTGetArea3
=== RUN   TestTDTFunc3/TDTGetArea4
    calc_test.go:153: [TDTGetArea4], want 11320.3200, but result is 5660.1600
        tearDown execute for TestTDTFunc3
--- FAIL: TestTDTFunc3 (0.00s)
    --- PASS: TestTDTFunc3/TDTGetArea1 (0.00s)
    --- FAIL: TestTDTFunc3/TDTGetArea2 (0.00s)
    --- PASS: TestTDTFunc3/TDTGetArea3 (0.00s)
    --- FAIL: TestTDTFunc3/TDTGetArea4 (0.00s)
FAIL
*/

func TestTDTFunc3(t *testing.T) {
	t.Cleanup(setUp(t.Name()))

	var tests = []struct {
		name string
		w, h float64
	}{
		{"TDTGetArea1", 10.1, 20.1},
		{"TDTGetArea2", 30.2, 40.2},
		{"TDTGetArea3", 50.3, 60.3},
		{"TDTGetArea4", 70.4, 80.4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r = GetArea(tt.w, tt.h)
			var wantR = tt.w * tt.h

			// var wantR float64
			// if ind%2 == 0 {
			// 	wantR = tt.w * tt.h
			// } else {
			// 	wantR = tt.w * tt.h * 2
			// }

			if r != wantR {
				t.Errorf("[%s], want %.4f, but result is %.4f\n", tt.name, wantR, r)
			}
		})
	}
}

func TestAutomic(t *testing.T) {
	var sum int32 = 0
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			// sum++
			atomic.AddInt32(&sum, 1) // 原子操作
		}()
	}
	wg.Wait()
	fmt.Println("\tsum =", sum) // ?
}

// // 性能(压力)测试
// func BenchmarkGetArea(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		GetArea(40, 50)
// 	}
// }

// // 性能(压力)测试
// func BenchmarkGetArea2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		GetArea(10, 20)
// 	}
// }

/*
go test -bench .
...
goos: windows
goarch: amd64
pkg: github.com/luoleiself/learning-go/testing
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
BenchmarkSuite/BenchmarkGetArea-8               1000000000               0.2637 ns/op
BenchmarkSuite/BenchmarkGetArea2-8              1000000000               0.2624 ns/op
        tearDown execute for BenchmarkSuite
PASS
package tearDown execute for TestMain
*/

func BenchmarkSuite(b *testing.B) {
	b.Cleanup(setUp(b.Name()))
	b.Run("BenchmarkGetArea", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetArea(10, 20)
		}
	})
	b.Run("BenchmarkGetArea2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetArea(40, 50)
		}
	})
}

// *testing.PB
// A PB is used by RunParallel for running parallel benchmarks.
// RunParallel 使用 PB 来运行并行基准测试
// func BenchMarkD(b *testing.B){
// 	b.RunParallel(func (pb *testing.PB){
// 		var buf bytes.Buffer
// 		for pb.Next() {
// 			buf.Reset()
// 		}
// 	})
// }
