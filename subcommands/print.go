package subcommands

import (
	"context"
	"flag"

	"github.com/dot96gal/subcommands-bubbletea-sample/bubbletea"
	gsubcmds "github.com/google/subcommands"
)

var _ gsubcmds.Command = (*PrintCommand)(nil)

type PrintCommand struct {
	delay int
}

func NewPrintCommand() *PrintCommand {
	return &PrintCommand{}
}

func (c *PrintCommand) Name() string {
	return "print"
}

func (c *PrintCommand) Synopsis() string {
	return "Print args to stdout."
}

func (c *PrintCommand) Usage() string {
	return `print [-delay] <some text>:
  Print args to stdout.
`
}

func (c *PrintCommand) SetFlags(f *flag.FlagSet) {
	f.IntVar(&c.delay, "delay", 0, "delay millisecond")
}

func (c *PrintCommand) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) gsubcmds.ExitStatus {
	p := bubbletea.NewPrintProgram(ctx, c.delay, f.Args())
	_, err := p.Run()
	if err != nil {
		return gsubcmds.ExitFailure
	}

	return gsubcmds.ExitSuccess
}
