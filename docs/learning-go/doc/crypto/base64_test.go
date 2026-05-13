package main

import (
	"encoding/base64"
	"testing"
)

func TestBase64(t *testing.T) {
	t.Run("base64 string", func(t *testing.T) {
		// 测试 base64 编码
		encoded := base64.StdEncoding.EncodeToString([]byte("hello 世界!"))
		t.Logf("EncodeToString: %s", encoded)

		// 测试 base64 解码
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			t.Fatalf("failed to decode: %v", err)
		}
		t.Logf("DecodeString: %s", decoded)
	})
	t.Run("Base64 []byte", func(t *testing.T) {
		// 测试 base64 编码
		encoded := make([]byte, base64.StdEncoding.EncodedLen(len("hello 世界!")))
		base64.StdEncoding.Encode(encoded, []byte("hello 世界!"))
		t.Logf("Encoded: %s", encoded)

		// 测试 base64 解码
		decoded := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
		n, err := base64.StdEncoding.Decode(decoded, encoded)
		if err != nil {
			t.Fatalf("failed to decode: %v", err)
		}
		t.Logf("Decoded: %s", decoded[:n])
	})
	t.Run("base64url", func(t *testing.T) {
		// 测试 base64url 编码
		encoded := base64.URLEncoding.EncodeToString([]byte("hello 世界!"))
		t.Logf("Encoded: %s", encoded)

		// 测试 base64url 解码
		decoded, err := base64.URLEncoding.DecodeString(encoded)
		if err != nil {
			t.Fatalf("failed to decode: %v", err)
		}
		t.Logf("decoded: %s", decoded)
	})
}
