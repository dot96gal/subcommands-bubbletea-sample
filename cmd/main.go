package main

import (
	"context"
	"flag"
	"os"

	"github.com/dot96gal/subcommands-bubbletea-sample/subcommands"
	gsubcmds "github.com/google/subcommands"
)

func main() {
	gsubcmds.Register(gsubcmds.HelpCommand(), "help")
	gsubcmds.Register(gsubcmds.FlagsCommand(), "help")
	gsubcmds.Register(gsubcmds.CommandsCommand(), "help")

	gsubcmds.Register(subcommands.NewPrintCommand(), "main")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(gsubcmds.Execute(ctx)))
}
