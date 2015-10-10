package main

import (
	"fmt"
	"flag"
	"math/rand"
	"time"
	"os"

	"github.com/dallasmarlow/composers"
)

var (
	listenAddr = flag.String(`addr`, `:8080`, `composer listener address`)
	composerName = flag.String(`name`, ``, `composer name`)
	randComposer = flag.Bool(`rand`, false, `random composer`)
)


func main() {
	flag.Parse()

	var composer composers.Composer
	if *composerName == `` || *randComposer {
		rand.Seed(time.Now().UnixNano())
		composer = composers.Rand()
	} else {
		composer = composers.Find(*composerName)
		if composer.Name == `` {
			fmt.Println(`unknown composer:`, *composerName)
			os.Exit(1)			
		}
	}

	if err := composer.Listen(*listenAddr); err != nil {
		fmt.Println(`composer exitied with err:`, err)		
	}
}