package khasher

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	kh := NewKHasher(5)
	hashes, err := kh.Hash([]byte("hello world"), 67)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hashes)
}
