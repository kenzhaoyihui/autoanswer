package main

import (
	"flag"
	"log"

	"os"

	ra "github.com/dracher/autoanswer/libs"
)

func main() {

	install := flag.Bool("i", false, "wether lanuch engine-setup")
	uninstall := flag.Bool("e", false, "wether lanuch engie-cleanup")

	log.Printf("install value is %d", *install)

	flag.Parse()

	log.Printf("current setup mode is: %t", *install)

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(-1)
	}

	ra.Run(true, true)
}
