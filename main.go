package main

import (
	"os"
	m "automate-cherrypicking/modules"
	c "automate-cherrypicking/modules/commits"

)
var provided_args []string

func main() {
	provided_args := os.Args[1:]

	if "postion" == provided_args[0] {
		c.Position(provided_args[1], provided_args[2])
	} else if "merge" == provided_args[0] {

	} else if "conflicts" == provided_args[0]{
		m.Conflicts(provided_args[1]);
	}
}


