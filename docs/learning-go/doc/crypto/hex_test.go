package main

import (
	"encoding/hex"
	"testing"
)

func TestHEX(t *testing.T) {
	t.Run("HEX string", func(t *testing.T) {
		encode := hex.EncodeToString([]byte("hello 世界!"))
		t.Logf("EncodeToString: %s", encode)

		deconde, err := hex.DecodeString(encode)
		if err != nil {
			t.Fatalf("failed to decode: %v", err)
		}
		t.Logf("DecodeString: %s", deconde)
	})
	t.Run("HEX []byte", func(t *testing.T) {
		encoded := make([]byte, hex.EncodedLen(len("hello 世界!")))
		hex.Encode(encoded, []byte("hello 世界!"))
		t.Logf("Encoded: %s", encoded)

		decoded := make([]byte, hex.DecodedLen(len(encoded)))
		n, err := hex.Decode(decoded, encoded)
		if err != nil {
			t.Fatalf("failed to decode: %v", err)
		}
		t.Logf("Decoded: %s", decoded[:n])
	})
}
