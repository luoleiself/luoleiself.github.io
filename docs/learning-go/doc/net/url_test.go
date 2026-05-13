package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestURL(t *testing.T) {
	t.Run("URL Parse", func(t *testing.T) {
		u, err := url.Parse("http://localhost:3000/sm1?name=zhangsan&age=18")
		if err != nil {
			t.Errorf("url.Parse err %v\n", err.Error())
		}
		t.Log(u)
	})
	t.Run("URL ParseRequestURI", func(t *testing.T) {
		buf := &bytes.Buffer{}
		buf.WriteString("name=lisi")
		buf.WriteString("age=22")
		buf.WriteString("addr=beijing")
		req, err := http.NewRequest(http.MethodPost, "http://localhost:3000", buf)
		if err != nil {
			t.Errorf("NewRequest err %v\n", err)
		}
		u, err := url.ParseRequestURI(req.URL.String())
		if err != nil {
			t.Errorf("ParseRequestURI err %v\n", err.Error())
		}
		t.Log(u)
	})
}
func TestURLQuery(t *testing.T) {
	t.Run("Query", func(t *testing.T) {
		v, err := url.ParseQuery("name=zhangsan&age=18&addr=beijing")
		if err != nil {
			t.Errorf("ParseQuery err %v\n", err.Error())
		}
		t.Log(v)
	})
}
func TestRetry(t *testing.T) {
	var Retry func(int, time.Duration, func() error) error

	t.Cleanup(func() func() {
		t.Log("cleanup setup...")
		Retry = func(attempt int, delay time.Duration, fn func() error) error {
			var err error
			for i := 0; i < attempt; i++ {
				if err = fn(); err == nil {
					return nil
				}
				t.Logf("Retry attempt %d failed: %v", i+1, err)
				time.Sleep(delay)
			}
			return fmt.Errorf("after %d attempt, last error: %v", attempt, err)
		}
		return func() {
			t.Log("cleanup after....")
			Retry = nil
		}
	}())
	t.Run("Retry", func(t *testing.T) {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		err := Retry(3, 1*time.Second, func() error {
			var v = random.Float32()
			if v > 0.5 {
				return fmt.Errorf("Retry failed: %v", v)
			}
			t.Logf("Retry success: %v", v)
			return nil
		})
		if err != nil {
			t.Errorf("Retry error: %v", err)
		} else {
			t.Log("Retry success")
		}
	})
}
