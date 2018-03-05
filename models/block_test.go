package models

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestGenerateBlock(t *testing.T) {
	b := GenerateBlock("0", "TEST_0", 0)
	fmt.Println(b)
}

func TestFormateBlock(t *testing.T) {
	b := GenerateBlock("0", "TEST_0", 0)
	jb, _ := json.Marshal(b)
	bf, err := FormateBlock(jb)
	if err != nil {
		fmt.Println(err)
	}
	if *bf != *b {
		t.Fail()
	}
	fmt.Println("bf == b")
}

func TestBlock_IsValid(t *testing.T) {
	b := GenerateBlock("0", "TEST_0", 0)
	b1 := GenerateBlock(b.Hash, "TEST_1", 1)
	b2 := GenerateBlock(b.Hash, "TEST_2", 1)
	fmt.Println("b1 is behind b:", b1.IsValid(b))
	fmt.Println("b2 is behind b1:", b2.IsValid(b1))
}