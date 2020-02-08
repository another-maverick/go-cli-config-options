package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
)

var flagOpts struct {
	Lang string `short:"l" long:"lang" default:"English" description:"select a language name"`
	Country string `short:"c" long:"country" default:"US" description:"Select a country"`
	Default bool `short:"d" long:"default" description:"Default output"`
}

func main() {
	flags.Parse(&flagOpts)

	if flagOpts.Default == true {
		fmt.Println("Default mode selected. You are boring!  Most of the US speaks English")
	} else {
		fmt.Printf("%s is widely spoken in %s \n", flagOpts.Lang, flagOpts.Country)
	}



}
