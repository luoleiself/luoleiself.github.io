package main

import (
	"os"
	"path"
	"testing"
)

func TestWriteViperConf(t *testing.T) {
	t.Cleanup(func() {
		os.Remove(path.Join("./", viperConfFile))
	})
	t.Run("writeViperConf", func(t *testing.T) {
		WriteViperConf()
	})
}
func TestReadViperConf(t *testing.T) {
	t.Cleanup(func() {
		os.Remove(path.Join("./", viperConfFile))
	})
	t.Run("ReadViperConf", func(t *testing.T) {
		file, err := os.Open(path.Join("./", viperConfFile))
		if os.IsNotExist(err) {
			t.Logf("\n%s file not exist\n", viperConfFile)
			WriteViperConf()
			ReadViperConf(nil)
		} else {
			defer file.Close()
			ReadViperConf(file)
		}
	})
}
func TestReadMyViperByBuffer(t *testing.T) {
	t.Run("ReadMyViperByBuffer", func(t *testing.T) {
		ReadMyViperByBuffer()
	})
}
