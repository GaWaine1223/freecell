// Copyright 2018 Lothar . All rights reserved.
// https://github.com/GaWaine1223

package models

import (
	"github.com/GaWaine1223/Lothar/freecell/common"
)

var theChain []*Block

func init() {
	theChain = make([]*Block, 0)
	Genesis := GenerateBlock([32]byte{0}, []byte("This is Genesis Block, Copyright Belong to Lothar"), 0)
	theChain = append(theChain, Genesis)
}

func Append(b *Block) error {
	if b.IsValid(GetTail()) {
		return common.Error(common.ErrInvalidBlock)
	}
	theChain = append(theChain, b)
	return nil
}

func GetTail() *Block {
	return theChain[GetLen()-1]
}

func GetLen() int {
	return len(theChain)
}