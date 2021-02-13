package main

import "flag"

var alphabet = flag.Int("alpha", 0, "Size of alphabet (words or chars)")
var length = flag.Int("len", 0, "Length of passphrase (words or chars)")
var entropy = flag.Float64("ent", 0, "Required entropy (bits)")

func main() {
	flag.Parse()
}
