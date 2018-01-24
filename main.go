package main

import (
	"fmt"

	"github.com/mushroomsir/meowcoin/pkg"
)

func main() {

	bc := pkg.NewBlockchain()

	bc.AddBlock("Reward 1 BTC to mushroom")
	bc.AddBlock("Rend 2 more BTC to mushroom")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
