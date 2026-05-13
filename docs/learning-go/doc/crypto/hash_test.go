package main

import (
	"crypto/md5"
	"crypto/sha256"
	"testing"
)

func TestMD5(t *testing.T) {
	t.Run("md5-1", func(t *testing.T) {
		t.Logf("md5.Sum(data) %x\n", md5.Sum([]byte("golang go go go!")))
		t.Logf("md5.BlockSize %d\n", md5.BlockSize)
		t.Logf("md5.Size %d\n", md5.Size)
	})
	t.Run("md5-2", func(t *testing.T) {
		m := md5.New()
		t.Logf("m.Sum(data) %x\n", m.Sum([]byte("hello world!")))
		t.Logf("m.Sum(data) %x\n", m.Sum([]byte("hello world!")))
		t.Logf("m.Sum(data) %x\n", m.Sum([]byte("gg!")))
		t.Logf("m.BlockSize() is %d\n", m.BlockSize())
		t.Logf("m.Size() is %v\n", m.Size())
	})
	t.Run("md5-3", func(t *testing.T) {
		m2 := md5.New()
		m2.Write([]byte("你好, 世界"))
		t.Logf("m2.Write(data) %x\n", m2.Sum(nil))
	})
}

func TestSHA256(t *testing.T) {
	t.Run("sha256-1", func(t *testing.T) {
		t.Logf("sha256.Sum256(data) %x\n", sha256.Sum256([]byte("golang go go go!")))
		t.Logf("sha256.BlockSize %d\n", sha256.BlockSize)
		t.Logf("sha256.Size %d\n", sha256.Size)
	})
	t.Run("sha256-2", func(t *testing.T) {
		s := sha256.New()
		t.Logf("s.Sum(data) %x\n", s.Sum([]byte("hello world!")))
		t.Logf("s.Sum(data) %x\n", s.Sum([]byte("hello world!")))
		t.Logf("s.Sum(data) %x\n", s.Sum([]byte("hello gg!")))
		t.Logf("s.BlockSize() is %d\n", s.BlockSize())
		t.Logf("s.Size() is %d\n", s.Size())
	})
	t.Run("sha256-3", func(t *testing.T) {
		s2 := sha256.New()
		s2.Write([]byte("你好, 世界"))
		t.Logf("s2.Write(data) %x\n", s2.Sum(nil))
	})
}
