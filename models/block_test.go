package models

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestGenerateBlock(t *testing.T) {
	b := GenerateBlock([32]byte{0}, []byte("TEST_0"), 0)
	fmt.Println(b)
}

func TestFormateBlock(t *testing.T) {
	b := GenerateBlock([32]byte{0}, []byte("TEST_0"), 0)
	jb, _ := json.Marshal(b)
	b2, err := FormateBlock(jb)
	if err != nil {
		fmt.Println(err)
	}
	if b2 != b {
		t.Fail()
	}
	fmt.Println((b2 == b))
}

func TestBlock_IsValid(t *testing.T) {
	b := GenerateBlock([32]byte{0}, []byte("TEST_0"), 0)
	b1 := GenerateBlock(b.Hash, []byte("TEST_1"), 1)
	fmt.Println(b1.IsValid(b))
}
