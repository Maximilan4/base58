package main

import (
	"fmt"
	"log"

	"github.com/Maximilan4/base58"
)

func main() {
	src := []byte("09")
	dst := make([]byte, 10)
	n, err := base58.EncodeBytes(src, dst)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v(%s)\n", dst[:n], dst[:n])

	// or encode string
	trg, err := base58.EncodeString("09")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(trg)
}
