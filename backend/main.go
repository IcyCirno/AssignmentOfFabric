package main

import (
	"blockchain/cmd"
)

func main() {

	cmd.Start()
	go cmd.Update()
}
