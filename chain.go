package main

import (
  "log"

  "github.com/dgraph-io/badger"
)

type Chain struct {
  // blocks []*Block
  tip []byte
  db *badger.DB
}

func (c *Chain) AddBlock(data string) {
  // prevBlock := c.blocks[len(c.blocks)-1]
  // newBlock := NewBlock(data, prevBlock.Hash)
  // c.blocks = append(c.blocks, newBlock)
  var (
    err error
    lastHash []byte
  )

  err = c.db.View(func(txn *badger.Txn) error {
    var item *badger.Item
    item, err = txn.Get([]byte("1"))
    if err != nil {
      return err
    }
    lastHash, err = item.Value()
    if err != nil {
      return err
    }
    return nil
  })

  newBlock := NewBlock(data, lastHash)

  err = c.db.Update(func (txn *badger.Txn) error {
    var err error
    err = txn.Set(newBlock.Hash, newBlock.Serialize())
    err = txn.Set([]byte("1"), newBlock.Hash)
    c.tip = newBlock.Hash
    return err
  })

  if err != nil {
    log.Panic(err)
  }
}

func NewGenesisBlock() *Block {
  return NewBlock("Genesis block", []byte{})
}

func NewChain() *Chain {
  // return &Chain{[]*Block{NewGenesisBlock()}}
  var tip []byte
  opts := badger.DefaultOptions
  opts.Dir = "./db"
  opts.ValueDir = "./db"
  db, err := badger.Open(opts)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  err = db.Update(func (txn *badger.Txn) error {
    var (
      err error
      item *badger.Item
    )
    item, err = txn.Get([]byte("1"))
    if err != nil {
      if err == badger.ErrKeyNotFound {
        genesis := NewGenesisBlock()
        err = txn.Set(genesis.Hash, genesis.Serialize())
        err = txn.Set([]byte("1"), genesis.Hash)
        return err
      }
      return err
    }
    tip, err = item.Value()
    return err
  })

  if err != nil {
    log.Panic(err)
  }

  // chain := Chain{genesis.Hash}
  chain := Chain{tip, db}
  return &chain
}
