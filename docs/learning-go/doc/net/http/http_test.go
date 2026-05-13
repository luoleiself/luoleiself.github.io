package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("sm1 GET", func(t *testing.T) {
		c := &http.Client{}
		res, err := c.Get("http://localhost:3000/sm1?name=sm1&age=1")
		if err != nil {
			t.Fatalf("GET error %v\n", err.Error())
		}
		res.Body.Close()
		t.Log(res.StatusCode)
	})
	t.Run("sm1 GET with params", func(t *testing.T) {
		params := url.Values{}
		u, err := url.Parse("http://localhost:3000/sm1")
		if err != nil {
			t.Fatalf("err %v\n", err.Error())
		}

		params.Set("name", "zhangsan")
		params.Set("age", "18")
		u.RawQuery = params.Encode()

		res, err := http.Get(u.String())
		if err != nil {
			t.Fatalf("GET sm1 with params err %v\n", err.Error())
		}
		defer res.Body.Close()
		t.Log(res.StatusCode)
	})
}

func TestPost(t *testing.T) {
	t.Run("sm2 POST with params", func(t *testing.T) {
		res, err := http.Post("http://localhost:3000/sm2?name=zhangsan&age=18", "application/x-www-form-urlencoded", strings.NewReader(""))
		if err != nil {
			t.Fatalf("GET sm1 with params err %v\n", err.Error())
		}
		defer res.Body.Close()
		t.Log(res.StatusCode)
	})
	t.Run("sm2 POST with request Body", func(t *testing.T) {
		params := url.Values{}
		params.Set("name", "wangwu")
		params.Set("age", "22")
		params.Set("addr", "guangzhou")

		formData := params.Encode()

		res, err := http.Post("http://localhost:3000/sm2", "application/x-www-form-urlencoded", strings.NewReader(formData))
		if err != nil {
			t.Fatalf("GET sm1 with params err %v\n", err.Error())
		}
		defer res.Body.Close()
		t.Log(res.StatusCode)
	})
	t.Run("sm2 POSTForm", func(t *testing.T) {
		res, err := http.PostForm("http://localhost:3000/sm2", url.Values{"name": {"lisi"}, "age": {"22"}, "addr": {"shanghai"}})
		if err != nil {
			t.Fatalf("POSTForm sm2 err %v\n", err.Error())
		}
		t.Log(res.StatusCode)
	})
}

func TestReqJSON(t *testing.T) {
	// curl --header 'Accept:application/json' http://localhost:3000/sm4
	t.Run("sm4 'Accept:application/json'", func(t *testing.T) {
		b := &bytes.Buffer{}
		req, err := http.NewRequest(http.MethodGet, "http://localhost:3000/sm4", b)
		if err != nil {
			t.Fatalf("NewRequest error %v\n", err)
		}
		req.Header.Set("Accept", "application/json")
		c := http.Client{}
		res, err := c.Do(req)
		if err != nil {
			t.Fatalf("client.Do error %v\n", err)
		}
		// defer res.Body.Close()
		len, _ := strconv.Atoi(res.Header.Get("Content-Length"))
		buf := make([]byte, len)
		// var buf []byte
		blen, err := res.Body.Read(buf)
		if err != nil && err != io.EOF {
			t.Fatalf("res.Body.Read error %v\n", err)
		}
		t.Logf("len is %d, data is %s\n", blen, buf)
	})
}

func TestPostForm(t *testing.T) {
	t.Run("sm3 POST multipart/form-data", func(t *testing.T) {
		buf := &bytes.Buffer{}

		body := multipart.NewWriter(buf)

		body.WriteField("field1", "value1") // 实际上调用 body.CreateFormFile() 写入指定值
		body.WriteField("field2", "value2")
		fWriter, err := body.CreateFormFile("file1", "file1.txt")
		if err != nil {
			t.Errorf("createFormFile err %v\n", err.Error())
		}
		fWriter.Write([]byte("file1 content"))
		fWriter2, err := body.CreateFormFile("file2", "file2.js")
		if err != nil {
			t.Errorf("CreateFormFile err %v\n", err.Error())
		}
		fWriter2.Write([]byte("console.log(\"hello world\")"))
		body.Close() // 关闭表单文件写入, 并将 trailing boundary 写入最后一行
		body.WriteField("boundary", body.Boundary())
		body.WriteField("formDataContentType", body.FormDataContentType())
		// t.Logf("表单文件分界符 %v \n", body.Boundary())
		// t.Logf("表单内容类型 %v\n", body.FormDataContentType())

		req, err := http.NewRequest(http.MethodPost, "http://localhost:3000/sm3", buf)
		req.Header.Set("Content-Type", body.FormDataContentType())
		if err != nil {
			t.Fatalf("POST multipart/form-data sm3 err %v\n", err.Error())
		}

		c := http.Client{}
		res, err := c.Do(req)
		if err != nil {
			t.Fatalf("POST multipart/form-data sm3 err %v\n", err.Error())
		}
		t.Log(res.StatusCode)
	})
}
