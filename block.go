package main

import (
	"crypto/sha256"
	"bytes"
  "strconv"
  "time"
)

type Block struct {
  Timestamp   int64
  Payload   []byte
  PrevBlockHash   []byte
  Hash  []byte
  Nonce int
}

func (b *Block) SetHash() {
  timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
  headers := bytes.Join([][]byte{b.PrevBlockHash, b.Payload, timestamp}, []byte{})
  hash := sha256.Sum256(headers)

  b.Hash = hash[:]
}

func NewBlock(data string, PrevBlockHash []byte) *Block {
  block := &Block{time.Now().Unix(), []byte(data), PrevBlockHash, []byte{}, 0}
  // block.SetHash()
  pow := NewProofOfWork(block)
  nonce, hash := pow.Run()

  block.Hash = hash[:]
  block.Nonce = nonce

  return block
}
