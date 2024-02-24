package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/scrot/musclemem-cli/cli"
	command "github.com/scrot/musclemem-cli/command/root"
)

var (
	name    = "musclemem"
	version = "0.0.1"
	author  = "Roy de Wildt"
	date    = ""
)

// TODO: make batching parallel
// TODO: add context handling
// TODO: check user logged in for some commands!
// TODO: init configuration file (baseurl, configpath)
// TODO: list single exercise / workout using wi/ei?
// TODO: make it possible to change config path
// TODO: create a test for each command
// TODO: see if build variables are loaded correctly
func main() {
	code := mainRun()
	os.Exit(int(code))
}

func mainRun() cli.ExitCode {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	config, err := cli.NewCLIConfig(name, version, author, date)
	if err != nil {
		fmt.Println(err)
		return cli.ExitError
	}

	root := command.NewRootCmd(config)
	if err := root.ExecuteContext(ctx); err != nil {
		switch {
		case errors.Is(err, cli.ErrNotAuthenticated):
			fmt.Println(err)
			return cli.ExitOK
		case errors.Is(err, cli.ErrExists):
			fmt.Println(err)
			return cli.ExitOK
		default:
			fmt.Println(err)
			return cli.ExitError
		}
	}

	return cli.ExitOK
}
