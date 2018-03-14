package main

import (
  "log"

	"github.com/dgraph-io/badger"
)

type ChainIterator struct {
  currentHash []byte
  db *badger.DB
}

func (c *Chain) Iterator() *ChainIterator {
  iter := &ChainIterator{c.tip, c.db}
  return iter
}

func (iter *ChainIterator) Next() *Block {
  var block *Block

  err := iter.db.View(func(txn *badger.Txn) error {
    // opts := badger.DefaultIteratorOptions
    // opts.PrefetchSize = 10
    // it := txn.NewIterator(opts)
    // defer it.Close()
    // for it.Rewind(); it.Valid(); it.Next() {
    //   item := it.Item()
    //   encodedBlock := item.Value()
    // }
    var (
      err error
      encodedBlock []byte
      item *badger.Item
    )
    item, err = txn.Get(iter.currentHash)
    if err != nil {
      return err
    }
    encodedBlock, err = item.Value()
    if err != nil {
      return err
    }
    block = DeserializeBlock(encodedBlock)
    return nil
  })

  if err != nil {
    log.Panic(err)
  }

  iter.currentHash = block.PrevBlockHash
  return block
}
