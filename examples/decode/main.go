package main

import (
	"fmt"
	"log"

	"github.com/Maximilan4/base58"
)

func main() {
	src := []byte("4ER")
	dst := make([]byte, 10)
	n, err := base58.DecodeBytes(src, dst)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dst[:n])

	// or encode string
	trg, err := base58.DecodeString("4ER")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(trg)
}
