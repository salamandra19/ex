package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

var f = flag.Int("algo", 256, "SHA алгоритм, допустимые значения 256, 384, 512")

func main() {
	flag.Parse()
	var h hash.Hash
	switch *f {
	case 256:
		h = sha256.New()
	case 384:
		h = sha512.New384()
	case 512:
		h = sha512.New()
	default:
		log.Fatal("incorrect algo, expect 256/384/512")
	}
	if _, err := io.Copy(h, os.Stdin); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", h.Sum(nil))
}
