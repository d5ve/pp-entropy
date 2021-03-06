package main

/*
pp-entropy

Given two of alphabet size, passphrase length, and bits of entropy,
calculate the third.

* entropy       = log2 ( alphabet size ) * length
* length        = entropy / log2( alphabet size )
* alphabet size = 2 ** ( entropy / length )

*/

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
		fmt.Printf("An alphabet size of %d and a passphrase length of %d gives entropy of %0.1f bits\n",
			*alphabet,
			*length,
			math.Log2(float64(*alphabet))*float64(*length))
	} else if *alphabet > 0 && *entropy > 0.0 && *length == 0 {
		fmt.Printf("An alphabet size of %d and an entropy of %0.1f bits requires a passphrase length of %d\n",
			*alphabet,
			*entropy,
			int(math.Ceil(*entropy/math.Log2(float64(*alphabet)))))
	} else if *length > 0 && *entropy > 0.0 && *alphabet == 0.0 {
		fmt.Printf("An entropy of %0.1f bits and a passphrase length of %d requires an alphabet size of %d\n",
			*entropy,
			*length,
			int(math.Ceil(math.Pow(2, *entropy/float64(*length)))))
	} else {
		flag.Usage()
	}
}

/*
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <http://unlicense.org/>
*/
