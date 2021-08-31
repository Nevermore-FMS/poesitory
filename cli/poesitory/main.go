package main

import (
	"os"

	"github.com/Nevermore-FMS/poesitory/cli/poesitory/cmd"
	"github.com/Nevermore-FMS/poesitory/cli/poesitory/eaobird"
)

func main() {
	if os.Getenv("POESITORY_CLI_GEN_DOCS") == "true" {
		cmd.GenDocs()
	}
	cmd.Execute()
}

func init() {
	eaobird.Print()
}
