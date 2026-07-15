package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/awerner/cli/internal"
)

func main() {
	fmt.Println(internal.Echo(strings.Join(os.Args[1:], " ")))
}
