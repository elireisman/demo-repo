package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core"
)

// utilize vulnerable import
func main() {
	bc := new(core.BlockChain)
	fmt.Printf("%+v\n", bc)
}
