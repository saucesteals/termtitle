package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/saucesteals/termtitle"
)

func main() {

	if err := termtitle.SetTitle(""); err != nil {
		panic(err)
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
