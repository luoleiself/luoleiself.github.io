package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"testing"
)

func TestJSONSlice(t *testing.T) {
	t.Run("slice -> json 1", func(t *testing.T) {
		var sl1 []int
		t.Logf("sl1 %v\n", sl1)
		jsonSl1, err := json.Marshal(sl1)
		if err != nil {
			t.Fatal("sl1 error", err)
		}
		// sl1 success null
		t.Logf("sl1 success %s\n", jsonSl1)
	})
	t.Run("slice -> json 2", func(t *testing.T) {
		var sl2 = []string{"hello world", "hello china", "hello beijing"}
		t.Logf("sl2 %v\n", sl2)
		jsonSl2, err := json.Marshal(sl2)
		if err != nil {
			t.Fatal("sl2 error", err)
		}
		// sl2 success ["hello world","hello china","hello beijing"]
		t.Logf("sl2 success %s\n", jsonSl2)
	})
}

func TestJSONMap(t *testing.T) {
	t.Run("map -> json 1", func(t *testing.T) {
		var m1 map[int]string
		t.Logf("m1 %v\n", m1)
		jsonM1, err := json.Marshal(m1)
		if err != nil {
			t.Fatal("m1 error", err)
		}
		// m1 success null
		t.Logf("m1 success %s\n", jsonM1)
	})
	t.Run("map -> json 2", func(t *testing.T) {
		var m2 = map[int]string{65: "A", 66: "B", 67: "C", 68: "D"}
		t.Logf("m2 %v\n", m2)
		jsonM2, err := json.Marshal(m2)
		if err != nil {
			t.Fatal("m2 error", err)
		}
		// m2 success {"65":"A","66":"B","67":"C","68":"D"}
		t.Logf("m2 success %s\n", jsonM2)
	})
	t.Run("map -> json 3", func(t *testing.T) {
		var m3 = map[string]int{"h": 104, "i": 105, "j": 106, "k": 107, "l": 108}
		t.Logf("m3 %v\n", m3)
		jsonM3, err := json.Marshal(m3)
		if err != nil {
			t.Fatal("m3 error", err)
		}
		// m3 success {"h":104,"i":105,"j":106,"k":107,"l":108}
		t.Logf("m3 success %s\n", jsonM3)
	})
}

func TestJSONStruct(t *testing.T) {
	t.Run("struct -> json 1", func(t *testing.T) {
		var stu1 Student
		// stu1 {Name: Sex:0 Age:0 Address: Weight:0 Hobbies:[] Score:map[] Avatar:<nil>}
		t.Logf("stu1 %+v\n", stu1)
		jsonStu1, err := json.Marshal(stu1)
		if err != nil {
			t.Fatal("stu1 error", err)
		}
		// stu1 success {"name":"","sex":0,"age":0,"addr":"","weight":0,"hobbies":null,"score":null,"avatar":null}
		t.Logf("stu1 success %s\n", jsonStu1)
		t.Logf("jsonStu1 is JSON %t\n", json.Valid(jsonStu1))
	})
	t.Run("struct -> json 2", func(t *testing.T) {
		var stu2 = Student{
			"zhangsan",
			1,
			18,
			"beijing china",
			95.2,
			[]string{"dance", "sing", "swim"},
			map[string]float32{"english": 98.5, "chinese": 100, "math": 99.6, "physics": 93},
			"binary file",
			"<strong>HomePage ZhangSan</strong>",
		}
		t.Logf("stu2 %+v\n", stu2)
		jsonStu2, err := json.Marshal(stu2)
		if err != nil {
			t.Fatal("stu2 error", err)
		}
		// stu2 success {"name":"zhangsan","sex":1,"age":18,"addr":"beijing china","weight":95.2,"hobbies":["dance","sing","swim"],"score":{"chinese":100,"english":98.5,"math":99.6,"physics":93}}
		t.Logf("stu2 success %s\n", jsonStu2)
		t.Logf("jsonStu2 is JSON %t\n", json.Valid(jsonStu2))
	})
}

func TestHTMLEscape(t *testing.T) {
	t.Run("HTMLEscape", func(t *testing.T) {
		var out bytes.Buffer
		json.HTMLEscape(&out, []byte(`{"Name": "<b>HTML content</b>"}`))
		out.WriteTo(os.Stdout)
		t.Log("\n")
	})
}

func TestEncoder(t *testing.T) {
	t.Run("encoder", func(t *testing.T) {
		var b = bytes.Buffer{}
		var encoder = json.NewEncoder(&b)
		encoder.Encode([]int{1, 2, 3, 4})
		encoder.SetEscapeHTML(true)
		encoder.Encode(map[int]string{1: "2", 2: "2", 3: "3", 4: "4"})
		t.Logf("%s\n", b.String())
		t.Log("--------")
		var decoder = json.NewDecoder(&b)
		t.Logf("decoder is more %t\n", decoder.More())
		t.Log(decoder.Buffered())
		for {
			token, err := decoder.Token()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			t.Logf("%T: %v", token, token)
			if decoder.More() {
				t.Log(" (more)")
			}
			t.Log("\n")
		}
	})
}
