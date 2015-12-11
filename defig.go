package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

const (
	figStart = `<figure`
	figEnd   = `</figure>`
	outFmt   = `
{{%% figure src="/post/%s/$1.jpg" alt="$2" caption="$3" %%}}
`
)

var (
	figRE      = regexp.MustCompile(`<figure.*?>.*?<img.*?src=".*?([-\w]+)-\d\d\dx\d\d\d.*?\.jpg.*?alt="(.*?)".*?<figcaption.*?>(.*?)</figcaption></figure>`)
	figStartRE = regexp.MustCompile(figStart)
	figEndRE   = regexp.MustCompile(figEnd)
	fnameRE    = regexp.MustCompile(`(\d+).*`)
)

func process(filename string) error {
	n := fnameRE.FindStringSubmatch(filename)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	figs := figStartRE.FindAllIndex(b, -1)
	fige := figEndRE.FindAllIndex(b, -1)
	if got, want := len(figs), len(fige); got != want {
		return fmt.Errorf("unclosed tag [%d != %d]", got, want)
	}
	for i := range figs {
		for j, c := range b[figs[i][1]:fige[i][0]] {
			if c == '\n' || c == '\r' {
				b[j+figs[i][1]] = ' '
			}
		}
	}
	of := []byte(fmt.Sprintf(outFmt, n[1]))
	return ioutil.WriteFile(filename+".fixed", figRE.ReplaceAll(b, of), os.FileMode(0666))
}

func main() {
	for _, fn := range os.Args[1:] {
		if err := process(fn); err != nil {
			log.Println(err)
		}
	}
}
