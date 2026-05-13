package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func WithValueNote() {
	type favContextKey string

	var k = favContextKey("language")

	var c1 = context.WithValue(context.Background(), k, "Go")
	var c2 = context.WithValue(c1, favContextKey("color"), "blue")
	find := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Printf("ctx.Value(%v) found value: %v\n", k, v)
			return
		}
		fmt.Printf("ctx.Value(%v) not found value\n", k)
	}
	fmt.Println("var k = favContextKey(\"language\")")
	fmt.Println("var c1 = context.WithValue(context.Background(), k, \"Go\")", c1)
	// context.Background.WithValue(type main.favContextKey, val Go)

	fmt.Println("var c2 = context.WithValue(c1, favContextKey(\"color\"), \"blue\")", c2)
	// context.Background.WithValue(type main.favContextKey, val Go).WithValue(type main.favContextKey, val blue)
	time.Sleep(time.Second * 2)
	find(c2, favContextKey("language")) // Go
	find(c2, favContextKey("color"))    // blue
	find(c2, favContextKey("name"))     // key not found: name
	t, ok := c2.Deadline()
	fmt.Println("c2.Deadline()", t, ok)                 // 0001-01-01 00:00:00 +0000 UTC false
	fmt.Println("c2.Err()", c2.Err())                   // <nil>
	fmt.Println("context.Cause(c2)", context.Cause(c2)) // <nil>
	time.Sleep(time.Second * 2)
}

func WithTimeoutNote() {
	fmt.Println("WithTimeout 第二个参数表示时长")
	c1, cancel1 := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel1()

	fmt.Println("now:", time.Now().Format(time.RFC3339Nano)) // 2024-12-25T10:20:21.1401299+08:00
	fmt.Println("c1, cancel1 := context.WithTimeout(context.Background(), time.Second*3)")
	fmt.Println("c1", c1) // context.Background.WithDeadline(2024-12-25 10:20:24.1401299 +0800 CST m=+10.439396301 [2.9860476s])
	fmt.Println("---------")

	select {
	case <-time.After(100 * time.Second):
		cancel1()
		t, ok := c1.Deadline()
		fmt.Println("c1.Deadline()", t, ok)                 // 2024-12-25 10:20:24.1401299 +0800 CST m=+10.439396301 true
		fmt.Println("c1 overslept")                         // overslept
		fmt.Println("c1.Err()", c1.Err())                   // context canceled
		fmt.Println("context.Cause(c1)", context.Cause(c1)) // context canceled
	case <-c1.Done():
		t, ok := c1.Deadline()
		fmt.Println("c1.Deadline()", t, ok)                 // 2024-12-25 10:20:24.1401299 +0800 CST m=+10.439396301 true
		fmt.Println("c1.Err()", c1.Err())                   // context deadline exceeded
		fmt.Println("context.Cause(c1)", context.Cause(c1)) // context deadline exceeded
	}
	time.Sleep(time.Second * 2)
}
func WithTimeoutCauseNote() {
	fmt.Println("withTimeoutCause 行为类似于 WithTimeout, 第三个参数指定超时的原因")
	c1, cancel1 := context.WithTimeoutCause(context.Background(), time.Second*3, errors.New("errors.New() generate cause"))
	defer cancel1()

	fmt.Println("now:", time.Now().Format(time.RFC3339Nano)) // 2024-12-25T18:13:32.5279514+08:00
	fmt.Println("c1, cancel1 := context.WithTimeoutCause(context.Background(), time.Second*3, errors.New(\"errors.New() generate cause\"))")
	fmt.Println("c1", c1) // c1 context.Background.WithDeadline(2024-12-25 18:13:35.5279514 +0800 CST m=+3.881755501 [2.9798665s])
	fmt.Println("---------")

	select {
	case <-time.After(time.Second * 100):
		cancel1()
		t, ok := c1.Deadline()
		fmt.Println("c1.Deadline()", t, ok)                 // 2024-12-25 18:13:35.5279514 +0800 CST m=+3.881755501 true
		fmt.Println("c1 overslept")                         // overslept
		fmt.Println("c1.Err()", c1.Err())                   // context canceled
		fmt.Println("context.Cause(c1)", context.Cause(c1)) // context canceled
	case <-c1.Done():
		t, ok := c1.Deadline()
		fmt.Println("c1.Deadline()", t, ok)                 // 2024-12-25 18:13:35.5279514 +0800 CST m=+3.881755501 true
		fmt.Println("c1.Err()", c1.Err())                   // context deadline exceeded
		fmt.Println("context.Cause(c1)", context.Cause(c1)) // errors.New() generate cause
	}
	time.Sleep(time.Second * 2)
}

func WithDeadlineNote() {
	fmt.Println("WithDeadline 第二个参数表示时间点")
	c1, cancel1 := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	defer cancel1()

	fmt.Println("now:", time.Now().Format(time.RFC3339Nano)) // 2024-12-25T10:28:16.9355001+08:00
	fmt.Println("c1, cancel1 := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))")
	fmt.Println("c1", c1) // context.Background.WithDeadline(2024-12-25 10:28:19.9355001 +0800 CST m=+16.129330801 [2.9993619s])
	fmt.Println("---------")

	select {
	case <-time.After(100 * time.Second):
		cancel1() // context canceled
		t, ok := c1.Deadline()
		fmt.Println("c1.Deadline()", t.Local(), ok)         // 2024-12-25 10:28:19.9355001 +0800 CST true
		fmt.Println("c1 overslept")                         // overslept
		fmt.Println("c1.Err()", c1.Err())                   // context canceled
		fmt.Println("context.Cause(c1)", context.Cause(c1)) // context canceled
	case <-c1.Done():
		t, ok := c1.Deadline()
		fmt.Println("c1.Deadline()", t.Local(), ok)         // 2024-12-25 10:28:19.9355001 +0800 CST true
		fmt.Println("c1.Err()", c1.Err())                   // context deadline exceeded
		fmt.Println("context.Cause(c1)", context.Cause(c1)) // context deadline exceeded
	}
	time.Sleep(time.Second * 2)
}

func WithDeadlineCauseNote() {
	fmt.Println("WithDeadlineCause 行为类似于 WithDeadline, 第三个参数指定超时的原因")
	c1, cancel1 := context.WithDeadlineCause(context.Background(), time.Now().Add(time.Second*3), errors.New("errors.New() generate cause"))
	defer cancel1()

	fmt.Println("now:", time.Now().Format(time.RFC3339Nano)) // 2024-12-25T18:29:46.6482356+08:00
	fmt.Println("c1, cancel1 := context.WithDeadlineCause(context.Background(), time.Now().Add(time.Second*3), errors.New(\"errors.New() generate cause\"))")
	fmt.Println("c1", c1) // context.Background.WithDeadline(2024-12-25 18:29:49.6482356 +0800 CST m=+4.290614101 [2.9742055s])
	fmt.Println("---------")

	select {
	case <-time.After(time.Second * 100):
		cancel1()
		t, ok := c1.Deadline()
		fmt.Println("c1.Deadline()", t.Local(), ok)         // 2024-12-25 18:29:49.6482356 +0800 CST true
		fmt.Println("c1 overslept")                         // overslept
		fmt.Println("c1.Err()", c1.Err())                   // context canceled
		fmt.Println("context.Cause(c1)", context.Cause(c1)) // context canceled
	case <-c1.Done():
		t, ok := c1.Deadline()
		fmt.Println("c1.Deadline()", t.Local(), ok)         // 2024-12-25 18:29:49.6482356 +0800 CST true
		fmt.Println("c1.Err()", c1.Err())                   // context deadline exceeded
		fmt.Println("context.Cause(c1)", context.Cause(c1)) // errors.New() generate cause
	}
	time.Sleep(time.Second * 2)
}

func WithCancelNote() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			defer close(dst)
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	c1, cancel1 := context.WithCancel(context.Background())
	defer cancel1()
	fmt.Println("c1, cancel1 := context.WithCancel(context.Background())", c1)
	fmt.Println("c1", c1) // context.Background.WithCancel
	fmt.Println("---------")
	time.Sleep(time.Second * 2)

	for n := range gen(c1) {
		fmt.Println(n) // 1 2 3 4 5
		if n == 5 {
			break
		}
	}
	t, ok := c1.Deadline()
	fmt.Println("c1.Deadline()", t, ok)                 // 0001-01-01 00:00:00 +0000 UTC false
	fmt.Println("c1.Err()", c1.Err())                   // <nil>
	fmt.Println("context.Cause(c1)", context.Cause(c1)) // <nil>
	time.Sleep(time.Second * 2)
}

func WithCancelCauseNote() {
	c1, cancel1 := context.WithCancelCause(context.Background())
	fmt.Println("c1, cancel1 := context.WithCancelCause(context.Background())")
	fmt.Println("c1", c1) // context.Background.WithCancel
	fmt.Println("---------")
	time.Sleep(time.Second * 2)

	cancel1(errors.New("ya mie dai"))
	t, ok := c1.Deadline()
	fmt.Println("c1.Deadline()", t, ok)                 // 0001-01-01 00:00:00 +0000 UTC false
	fmt.Println("c1.Err()", c1.Err())                   // context canceled
	fmt.Println("context.Cause(c1)", context.Cause(c1)) // ya mie dai
	time.Sleep(time.Second * 2)
}
