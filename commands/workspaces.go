package commands

import (
	"fmt"

	"github.com/chinmaygupta28/asana/api"
	"github.com/urfave/cli"
)

func Workspaces(c *cli.Context) {
	for _, w := range api.Me().Workspaces {
		fmt.Printf("%16d %s\n", w.Id, w.Name)
	}
}
