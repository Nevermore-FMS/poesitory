package main

import (
	"github.com/Nevermore-FMS/poesitory/cli/poesitory/cmd"
	"github.com/Nevermore-FMS/poesitory/cli/poesitory/eaobird"
)

func main() {
	cmd.Execute()
}

func init() {
	eaobird.Print()
}
