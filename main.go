package main

// pp-entropy
// Given two of alphabet size, passphrase length, and bits of entropy,
// calculate the third.

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintln(os.Stderr, "pp-entropy - random passphrase strength calculator")
	fmt.Fprintln(os.Stderr, "Given two of alphabet size, passphrase length, and bits of entropy, calculate the third.")
	flag.PrintDefaults()

}
func main() {

	alphabet := flag.Int("alpha", 0, "Size of alphabet (words or chars)")
	length := flag.Int("len", 0, "Length of passphrase (words or chars)")
	entropy := flag.Float64("ent", 0, "Required entropy (bits)")

	flag.Usage = usage
	flag.Parse()
	fmt.Println("A=", *alphabet)
	fmt.Println("L=", *length)
	fmt.Println("E=", *entropy)

	if *entropy > 10.0 {
		fmt.Println("ent 10")
	}
	flag.Usage()
}
