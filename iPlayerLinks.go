package main

import (
	"flag"

	"github.com/gandalf15/iplayerlinks/cli"
	"github.com/gandalf15/iplayerlinks/gui"
)

func main() {
	urlPtr := flag.String("url", "", "-url=[iPlayer URL with episodes]")
	audioDescribedPtr := flag.Bool("audioDescribed", false, "-audioDescribed=[bool]")
	signLangPtr := flag.Bool("signLang", false, "-signLang=[bool]")
	flag.Parse()
	if *urlPtr == "" {
		gui.Gui()
	} else {
		cli.Cli(urlPtr, audioDescribedPtr, signLangPtr)
	}

}
