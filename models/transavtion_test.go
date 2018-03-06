package models

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestGenerateTransWithID(t *testing.T) {
	_, err := Login("luda")
	if err != nil {
		t.FailNow()
	}
	tr, err := GenerateTransWithID("luda", "TRANS_1")
	if err != nil {
		fmt.Println("generateTrans fail", err)
		t.FailNow()
	}
	err = tr.IsVaild()
	if err != nil {
		fmt.Println("invalid trans, err =", err)
	}
	fmt.Println(tr)
}

func TestGenerateTransWithKey(t *testing.T) {
	u, err := Login("yoghurt")
	if err != nil {
		t.FailNow()
	}
	tr, err := GenerateTransWithKey(u.Public, u.Private, "TRANS_2")
	if err != nil {
		fmt.Println("generateTrans fail", err)
		t.FailNow()
	}
	err = tr.IsVaild()
	if err != nil {
		fmt.Println("invalid trans, err =", err)
	}
	fmt.Println(tr)
}

func TestFormateTrans(t *testing.T) {
	_, err := Login("luda")
	if err != nil {
		t.FailNow()
	}
	tr, err := GenerateTransWithID("luda", "TRANS_1")
	if err != nil {
		fmt.Println("generateTrans fail", err)
		t.FailNow()
	}
	btr, _ := json.Marshal(tr)
	tr2, err := FormateTrans(btr)
	if err != nil {
		fmt.Println("formate fail", err)
		t.FailNow()
	}
	err = tr2.IsVaild()
	if err != nil {
		fmt.Println("invalid trans, err =", err)
	}
	fmt.Println(tr2)
}