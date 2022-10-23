// The `four` command randomly chooses four short words from
// `/usr/share/dict/words`.
package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
)

const wordsFile = "/usr/share/dict/words"

func main() {
	f, err := os.Open(wordsFile)
	if err != nil {
		log.Fatalf("Couldn't open words file: %v", err)
	}
	defer f.Close()
	var words []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		w := sc.Text()
		if len(w) > 6 {
			continue
		}
		words = append(words, w)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't scan words file: %v", err)
	}
	l := big.NewInt(int64(len(words)))
	for i := 0; i < 4; i++ {
		r, err := rand.Int(rand.Reader, l)
		if err != nil {
			log.Fatalf("Couldn't read random source: %v", err)
		}
		fmt.Printf("%s", words[int(r.Int64())])
	}
	fmt.Println()
}
