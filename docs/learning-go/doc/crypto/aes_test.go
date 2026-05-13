package main

import (
	"io"
	"os"
	"testing"
)

func TestAES(t *testing.T) {
	t.Run("AES-1", func(t *testing.T) {
		encoded, err := encrypt("")
		if err != nil {
			t.Fatalf("encoded failed: %v\n", err)
		}
		t.Logf("encoded: %v", encoded)

		decoded, err := decrypt(encoded)
		if err != nil {
			t.Fatalf("decoded failed: %v\n", err)
		}
		t.Logf("decoded: %v", decoded)
	})
	t.Run("AES-2", func(t *testing.T) {
		encoded, err := encrypt("hello 世界!")
		if err != nil {
			t.Fatalf("encoded failed: %v\n", err)
		}
		t.Logf("encoded: %v", encoded)

		decoded, err := decrypt(encoded)
		if err != nil {
			t.Fatalf("decoded failed: %v\n", err)
		}
		t.Logf("decoded: %v", decoded)
	})
	t.Run("AES-3", func(t *testing.T) {
		file, err := os.Open("./testdata/Person.xml")
		if err != nil {
			t.Fatalf("open file failed: %v\n", err)
		}
		defer file.Close()
		f, err := io.ReadAll(file)
		if err != nil {
			t.Fatalf("read file failed: %v\n", err)
		}

		encoded, err := encrypt(string(f))
		if err != nil {
			t.Fatalf("encoded failed: %v\n", err)
		}
		t.Logf("encoded: %v", encoded)

		decoded, err := decrypt(encoded)
		if err != nil {
			t.Fatalf("decoded failed: %v\n", err)
		}
		t.Logf("decoded: %v", decoded)
	})
}
