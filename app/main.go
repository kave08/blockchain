package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

}

func run() error {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return err
	}
	fmt.Println(privateKey)

	return nil
}
