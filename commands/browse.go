package commands

import (
	"os/exec"
	"strconv"

	"github.com/urfave/cli"

	"github.com/chinmaygupta28/asana/api"
	"github.com/chinmaygupta28/asana/config"
	"github.com/chinmaygupta28/asana/utils"
)

func Browse(c *cli.Context) {
	taskId := api.FindTaskId(c.Args().First(), true)
	url := "https://app.asana.com/0/" + strconv.Itoa(config.Load().Workspace) + "/" + taskId
	launcher, err := utils.BrowserLauncher()
	utils.Check(err)
	cmd := exec.Command(launcher, url)
	err = cmd.Start()
	utils.Check(err)
}
