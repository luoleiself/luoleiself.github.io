package main

import (
	"mime"
	"testing"
)

func TestMIME(t *testing.T) {
	t.Run("MIME AddExtensionType", func(t *testing.T) {
		t.Logf("设置以点开始的文件扩展名的 MIME 类型\n")
		err1 := mime.AddExtensionType(".txt", "text/plain")
		if err1 != nil {
			t.Errorf("AddExtensionType err1 err %v\n", err1.Error())
		}
		err2 := mime.AddExtensionType(".html", "text/html")
		if err2 != nil {
			t.Errorf("AddExtensionType err2 err %v\n", err2.Error())
		}
		err3 := mime.AddExtensionType(".js", "application/js")
		if err3 != nil {
			t.Errorf("AddExtensionType err3 err %v\n", err3.Error())
		}
	})
	t.Run("MIME FormatMediaType", func(t *testing.T) {
		t.Logf("返回指定 MIME 类型和给定参数序列化后的字符串\n")
		result1 := mime.FormatMediaType("text/html", map[string]string{"charset": "utf-8"})
		t.Logf("FormatMediaType %v\n", result1)
		result2 := mime.FormatMediaType("application/json", map[string]string{"charset": "gbk"})
		t.Logf("FormatMediaType %v\n", result2)
	})
	t.Run("MIME ParseMediaType", func(t *testing.T) {
		t.Logf("返回指定参数解析后的媒体类型和参数组成的 map, 媒体类型是 Headers 中 Content-Type 和 Content-Disposition 的值\n")
		m1, p1, err1 := mime.ParseMediaType("text/html; charset=utf-8")
		if err1 != nil {
			t.Errorf("ParseMediaType err1 err %v\n", err1.Error())
		}
		t.Logf("mediaType %v params %v \n", m1, p1)
		m2, p2, err2 := mime.ParseMediaType("application/json; charset=gbk")
		if err2 != nil {
			t.Errorf("ParseMediaType err2 err %v\n", err2.Error())
		}
		t.Logf("mediaType %v params %v \n", m2, p2)
	})
	t.Run("MIME ExtensionByType", func(t *testing.T) {
		t.Logf("返回 MIME 类型关联的以点开始的文件扩展名组成的 slice\n")
		result1, err1 := mime.ExtensionsByType("text/plain")
		if err1 != nil {
			t.Errorf("ExtensionByType err1 err %v\n", err1.Error())
		}
		t.Logf("ExtensionByType %v\n", result1)
		result2, err2 := mime.ExtensionsByType("application/pdf")
		if err2 != nil {
			t.Errorf("ExtensionByType err2 err %v\n", err2.Error())
		}
		t.Logf("ExtensionByType %v\n", result2)
		result3, err3 := mime.ExtensionsByType("audio/wav")
		if err3 != nil {
			t.Errorf("ExtensionByType err3 err %v\n", err3.Error())
		}
		t.Logf("ExtensionByType %v\n", result3)
	})
	t.Run("MIME TypeByExtension", func(t *testing.T) {
		t.Logf("返回以点开始的文件扩展名关联的 MIME 类型\n")
		result1 := mime.TypeByExtension(".webp")
		t.Logf("TypeByExtension %v\n", result1)
		result2 := mime.TypeByExtension(".wma")
		t.Logf("TypeByExtension %v\n", result2)
		result3 := mime.TypeByExtension(".xlsx")
		t.Logf("TypeByExtension %v\n", result3)
	})
}

func TestMIMECoder(t *testing.T) {
	t.Run("MIME WordDecoder", func(t *testing.T) {
		dec := mime.WordDecoder{}
		header, err := dec.Decode("=?utf-8?q?=C2=A1Hola,_se=C3=B1or!?=")
		if err != nil {
			t.Errorf("WordDecoder err err %v\n", err.Error())
		}
		t.Logf("WordDecoder %v\n", header)
	})
	t.Run("MIME WordEncoder", func(t *testing.T) {
		t.Logf("WordEncoder %v\n", mime.QEncoding.Encode("utf-8", "¡Hola, señor!"))
		t.Logf("WordEncoder %v\n", mime.QEncoding.Encode("utf-8", "Hello!"))
		t.Logf("WordEncoder %v\n", mime.BEncoding.Encode("UTF-8", "¡Hola, señor!"))
		t.Logf("WordEncoder %v\n", mime.QEncoding.Encode("ISO-8859-1", "Caf\xE9"))
	})
}
