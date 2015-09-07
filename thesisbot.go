package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/thoj/go-ircevent"
)

const (
	nick = "thesis_bot"
	rate = float64(10 * time.Minute)
)

func main() {
	rand.Seed(time.Now().UnixNano())
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, os.Kill)

	i := irc.IRC(nick, nick)
	if err := i.Connect("irc.freenode.net:6667"); err != nil {
		panic(err)
	}
	i.AddCallback("001", func(e *irc.Event) {
		i.Join("##maclab")
	})
	i.AddCallback("JOIN", func(e *irc.Event) {
		if e.Nick != nick {
			return
		}
		go func() {
			for {
				next := time.Duration(rand.ExpFloat64() * rate) 
				log.Printf("Next thesis after %v (%v)", next, time.Now().Add(next))
				<-time.After(next)
				i.Privmsg("##maclab", "thesis")
			}
		}()
	})
	fmt.Println("^C to quit")
	<-sigint
	i.Quit()
}
