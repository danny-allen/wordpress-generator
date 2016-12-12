package main

import (
	"os"
	"dao/wordpress/download"
	"dao/wordpress/create"
	"fmt"
)

// First contact.
func main() {

	// Check an argument exists.
	if(len(os.Args) > 1 && os.Args != nil) {

		// We have a command, lets see if it exists.
		tryCommand()
	} else {

		// We dont have
		fmt.Println("WordPress: Please specify a command.")
	}
}

// Run the command, if exists.
func tryCommand() {

	// Check for download param
	switch os.Args[1] {

	// CMD: wordpress download
	case "download":

		//Run download, accept no flags.
		download.Run()
		break

	// CMD: wordpress create
	case "create":

		create.Make(os.Args[2])
		break;
	default:
		fmt.Println("WordPress: Did not recognise that command.")
	}
}


