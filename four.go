// The four binary randomly chooses four words from /usr/share/dict/words.
package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

const wordsFile = "/usr/share/dict/words"

func main() {
	f, err := os.Open(wordsFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var words []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		words = append(words, sc.Text())
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	l := big.NewInt(int64(len(words)))
	for i := 0; i < 4; i++ {
		r, err := rand.Int(rand.Reader, l)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", words[int(r.Int64())])
	}
	fmt.Println()
}
