package main

import (
	"os"

	"github.com/awerner/cli/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
