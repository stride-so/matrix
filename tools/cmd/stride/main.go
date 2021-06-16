package stride

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type parseCommand struct {
	fs *flag.FlagSet
	in string
}

func newParseCommand() *parseCommand {
	pc := &parseCommand{
		fs: flag.NewFlagSet("parse", flag.ContinueOnError),
	}
	pc.fs.StringVar(&pc.in, "in", "", "filename to be parsed")
	return pc
}

func (c *parseCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *parseCommand) Run() error {
	fmt.Println("would parse", c.in)
	return nil
}

func (c *parseCommand) Name() string {
	return c.fs.Name()
}

type runner interface {
	Run() error
	Init([]string) error
	Name() string
}

func root(args []string) error {
	if len(args) < 1 {
		return errors.New("you must pass a command")
	}
	cmds := []runner{
		newParseCommand(),
	}

	subCmd := os.Args[1]
	for _, cmd := range cmds {
		if cmd.Name() == subCmd {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}
	return fmt.Errorf("unknown command: %s", subCmd)
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
