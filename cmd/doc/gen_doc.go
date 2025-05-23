package main

import (
	"log"

	"github.com/free5gc/cdrFileParser/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenMarkdownTree(cmd.RootCmd, "../../docs")
	if err != nil {
		log.Fatal(err)
	}
}
