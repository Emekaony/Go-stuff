package logs

import (
	"log"
)

func UsingFatal() {
	// set the prefix to be the name of the function
	log.SetPrefix("UsingFatal: ")
	// log.fatal ends the program
	log.Fatal("This will crash the program")
	log.Print("can u see me?")
}
