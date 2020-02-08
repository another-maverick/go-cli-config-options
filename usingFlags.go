package main

import (
	"flag"
	"fmt"
)

var language = flag.String("language", "English", "Which language do you want to write in..")

var langVar bool

func init() {
	flag.BoolVar(&langVar, "french", false, "select French Language" )
	flag.BoolVar(&langVar, "f", false, "select French Language" )
}

func main() {
	flag.Parse()

	if langVar == true {
		fmt.Println("You have selected french")
	}  else {
		fmt.Println("you have selected English")
	}

	//printing defaults
	flag.PrintDefaults()

	flag.VisitAll(func(flag *flag.Flag){
		fmt.Printf("Overriding the help string -- Name: %s, DefaultValue: %s, Usage: %s \n", flag.Name, flag.DefValue, flag.Usage)
	})
}
