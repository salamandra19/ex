package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var f = flag.Int("algo", 256, "SHA алгоритм, допустимые значения 256, 384, 512")

func main() {
	flag.Parse()
	for _, data := range flag.Args() {
		switch *f {
		case 256:
			fmt.Printf("%x\n", sha256.Sum256([]byte(data)))
		case 384:
			fmt.Printf("%x\n", sha512.Sum384([]byte(data)))
		case 512:
			fmt.Printf("%x\n", sha512.Sum512([]byte(data)))
		default:
			fmt.Println("incorrect algo, expect 256/384/512")
		}
	}
}
