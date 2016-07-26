package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var f384 = flag.Bool("384", false, "преобразование SHA384")
var f512 = flag.Bool("512", false, "преобразование SHA512")

func main() {
	flag.Parse()
	for _, data := range flag.Args() {
		fmt.Println(data)
		if !*f384 && !*f512 {
			fmt.Printf("%x\n", sha256.Sum256([]byte(data)))
		} else if *f384 {
			fmt.Printf("%x\n", sha512.Sum384([]byte(data)))
		} else {
			fmt.Printf("%x\n", sha512.Sum512([]byte(data)))
		}
	}

}
