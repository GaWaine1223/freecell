// Copyright 2018 Lothar . All rights reserved.
// https://github.com/GaWaine1223

package models

import (
	"strconv"
	"time"
	"fmt"
	"encoding/json"

	dhash "github.com/GaWaine1223/Lothar/sha256"

	"github.com/GaWaine1223/Lothar/freecell/common"
)

type Block struct {
	PVHash		string
	Timestamp	int64
	Data		string
	Index		int64
	Nonce		int64
	Hash 		string
}

// Formate received []byte to a block object
func FormateBlock(b []byte) (*Block, error) {
	block := &Block{}
	err := json.Unmarshal(b, block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// Generate a new block, it takes sometime and can be stopped by using the following function.
// hash = PVHash+Timestamp+Data+n+Nonce
func GenerateBlock(pvHash, data string, index int64) *Block {
	var metaData string
	time := time.Now().UnixNano()
	tStr := strconv.FormatInt(time, 10)
	nStr := strconv.FormatInt(index, 10)
	metaData = pvHash + tStr + data + nStr
	hash, nonce := dhash.HashwithDifficulty([]byte(metaData), common.HashDifficulty)
	return &Block{
		PVHash :pvHash,
		Timestamp:time,
		Data :data,
		Index: index,
		Nonce :nonce,
		Hash :fmt.Sprint(hash),
	}
}

func (b *Block) Interupt() bool {
	return dhash.StopHash()
}

func (b *Block) IsValid(pvb *Block) bool {
	var metaData string
	var byteHash [dhash.Size]byte
	if b.PVHash != pvb.Hash {
		return false
	}
	tStr := strconv.FormatInt(b.Timestamp, 10)
	nStr := strconv.FormatInt(b.Index, 10)
	noStr := strconv.FormatInt(b.Nonce, 10)
	metaData = b.PVHash + tStr + b.Data + nStr + noStr
	copy(byteHash[:], b.Hash)
	return dhash.Verification([]byte(metaData), byteHash)
}