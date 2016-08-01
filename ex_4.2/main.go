package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var f = flag.Int("algo", 256, "SHA алгоритм, допустимые значения 256, 384, 512")

func main() {
	flag.Parse()
	switch *f {
	case 256:
		h := sha256.New()
		if _, err := io.Copy(h, os.Stdin); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%x\n", h.Sum(nil))
	case 384:
		h := sha512.New384()
		if _, err := io.Copy(h, os.Stdin); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%x\n", h.Sum(nil))
	case 512:
		h := sha512.New()
		if _, err := io.Copy(h, os.Stdin); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%x\n", h.Sum(nil))
	default:
		fmt.Println("incorrect algo, expect 256/384/512")
	}
}
