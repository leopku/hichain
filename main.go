package main

import (
	"fmt"
)

func main()  {
  fmt.Println("Hello block chain!")
  c := NewChain()

  c.AddBlock("Block 1")
  c.AddBlock("Block 2")

  for _, block := range c.blocks {
    fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
    fmt.Printf("Payload: %s\n", block.Payload)
    fmt.Printf("Hash: %x\n", block.Hash)
    fmt.Println()
  }
}
