package main

import (
	"github.com/mushroomsir/meowcoin/cli"
	"github.com/mushroomsir/meowcoin/pkg"
)

func main() {

	bc := pkg.NewBlockchain()
	defer bc.CloseDB()

	cli := cli.NewCLI(bc)
	cli.Run()
}
