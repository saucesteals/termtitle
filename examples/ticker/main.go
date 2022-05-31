package main

import (
	"fmt"
	"time"

	"github.com/saucesteals/termtitle"
)

func main() {

	tick := time.NewTicker(time.Second)
	start := time.Now()

	for {
		select {
		case <-tick.C:
			if err := termtitle.SetTitle(fmt.Sprintf("Elapsed: %.0fs", time.Since(start).Seconds())); err != nil {
				panic(err)
			}
		}
	}

}
