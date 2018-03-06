// Copyright 2018 Lothar . All rights reserved.
// https://github.com/GaWaine1223

package models

import (
	"encoding/json"

	"github.com/GaWaine1223/Lothar/keygen"
)

type Trans struct {
	// public key
	Account		string		`json:"account"`
	Cipher		string        `json:"cipher"`
	Transaction	string        `json:"transaction"`
}

func (t *Trans) IsVaild() error {
	return keygen.Verify(t.Account, t.Cipher, []byte(t.Transaction))
}

func FormateTrans(b []byte) (*Trans, error) {
	trans := &Trans{}
	err := json.Unmarshal(b, trans)
	if err != nil {
		return nil, err
	}
	return trans, nil
}

func GenerateTransWithID(id, data string) (*Trans, error) {
	a, c, err := keygen.Signature(id, []byte(data))
	if err != nil {
		return nil, err
	}
	return &Trans{
		Account: a,
		Cipher: c,
		Transaction: data}, nil
}

func GenerateTransWithKey(pb, pv, data string) (*Trans, error) {
	c, err := keygen.Signature2(pv, []byte(data))
	if err != nil {
		return nil, err
	}
	return &Trans{
		Account: pb,
		Cipher: c,
		Transaction: data}, nil
}