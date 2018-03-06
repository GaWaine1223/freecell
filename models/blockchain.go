// Copyright 2018 Lothar . All rights reserved.
// https://github.com/GaWaine1223

package models

import (
	"encoding/json"

	"github.com/GaWaine1223/Lothar/freecell/common"
)

type TheChain struct {
	Chain []*Block        `json:"chain"`
}

var singleChain *TheChain

func init() {
	singleChain = newChain()
	Genesis := GenerateBlock("0", "This is Genesis Block, Copyright Belong to Lothar", 0)
	singleChain.Chain = append(singleChain.Chain, Genesis)
}

func newChain() *TheChain {
	theChain := make([]*Block, 0)
	return &TheChain{theChain}
}
func AppendChain(b *Block) error {
	if !b.IsValid(GetChainTail()) {
		return common.Error(common.ErrInvalidBlock)
	}
	singleChain.Chain = append(singleChain.Chain, b)
	return nil
}

func GetChainTail() *Block {
	return singleChain.Chain[GetChainLen()-1]
}

func GetChainLen() int64 {
	return int64(len(singleChain.Chain))
}

func ReplaceChain(c2 *TheChain) error {
	if int64(len(c2.Chain)) <= GetChainLen() {
		return common.Error(common.ErrInvalidBlock)
	}
	for i, b := range c2.Chain {
		if i == 0 {
			if *c2.Chain[i] != *singleChain.Chain[i] {
				return common.Error(common.ErrInvalidGenesisBlock)
			}
			continue
		}
		if !b.IsValid(c2.Chain[i-1]) {
			return common.Error(common.ErrInvalidBlock)
		}
	}
	singleChain.Chain = c2.Chain
	return nil
}

func FetchChain() *TheChain {
	return singleChain
}

func FormateChain(b []byte) (*TheChain, error) {
	c := &TheChain{}
	err := json.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}
	return c, err
}