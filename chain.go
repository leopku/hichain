package main

import (
  "github.com/dgraph-io/badger"
)

type Chain struct {
  blocks []*Block
}

func (c *Chain) AddBlock(data string) {
  prevBlock := c.blocks[len(c.blocks)-1]
  newBlock := NewBlock(data, prevBlock.Hash)
  c.blocks = append(c.blocks, newBlock)
}

func NewGenesisBlock() *Block {
  return NewBlock("Genesis block", []byte{})
}

func NewChain() *Chain {
  return &Chain{[]*Block{NewGenesisBlock()}}
}
