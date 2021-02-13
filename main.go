package main

// pp-entropy
// Given two of alphabet size, passphrase length, and bits of entropy,
// calculate the third.
// - entropy       = log2 ( alphabet size ) * length
// - length        = entropy / log2( alphabet size )
// - alphabet size = 2 ** entropy *
//                  x        = log2 ( a ** b )
//		    2 ** x   = a ** b
//                  (2 ** x) / log(b) = a ???
import (
	"flag"
	"fmt"
	"math"
	"os"
)

func usage() {
	fmt.Fprintln(os.Stderr, "pp-entropy - random passphrase strength calculator")
	fmt.Fprintln(os.Stderr, "Given two of alphabet size, passphrase length, and bits of entropy, calculate the third.")
	flag.PrintDefaults()

}
func main() {
	flag.Usage = usage
	alphabet := flag.Int("alpha", 0, "Size of alphabet (words or chars)")
	length := flag.Int("len", 0, "Length of passphrase (words or chars)")
	entropy := flag.Float64("ent", 0, "Required entropy (bits)")
	flag.Parse()

	if *alphabet > 0 && *length > 0 && *entropy == 0.0 {
		fmt.Printf("An alphabet of %d and a length of %d gives entropy of %0.1f bits\n",
			*alphabet,
			*length,
			math.Log2(float64(*alphabet))*float64(*length))
	} else if *alphabet > 0 && *entropy > 0.0 && *length == 0 {
		fmt.Printf("An alphabet of %d and an entropy of %0.1f bits requires a length of %d\n",
			*alphabet,
			*entropy,
			int(math.Ceil(*entropy/math.Log2(float64(*alphabet)))))
	} else if *length > 0 && *entropy > 0.0 && *alphabet == 0.0 {
		fmt.Printf("An entropy of %0.1f bits and a length = %d requires an alphabet size of %d\n",
			*entropy,
			*length,
			int(math.Ceil(math.Log(math.Pow(2, *entropy))/math.Log(float64(*length)))))
	}
}
