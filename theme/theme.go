package theme

import (
	"flag"
	"fmt"
)

func Scaffold() {

	// Declare flags.
	name := flag.String("name", "test", "The name of your theme.")

	// Command line parsing of flags.
	flag.Parse()

	// Output.
	fmt.Println("The theme you are making is called: " + *name)
}