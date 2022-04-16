package main

import (
	"woocoin/blockchain"
	"woocoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
