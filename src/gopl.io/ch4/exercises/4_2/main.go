package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	mode := flag.Int("size", 256, "Hash size: 256, 384, 512")
	flag.Parse()
	buf := make([]byte, 80)
	_, err := os.Stdin.Read(buf)
	if err != nil {
		return
	}
	switch *mode {
	case 256:
		res := sha256.Sum256(buf)
		fmt.Printf("Hash %d: %x\n", *mode, res)
	case 384:
		res := sha512.Sum384(buf)
		fmt.Printf("Hash %d: %x\n", *mode, res)
	case 512:
		res := sha512.Sum512(buf)
		fmt.Printf("Hash %d: %x\n", *mode, res)
	default:
		fmt.Printf("Mode %d not supported\n", *mode)
		os.Exit(1)
	}
}
