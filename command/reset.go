package command

import (
	"strings"
)

type ResetCommand struct {
	Meta
}

func (c *ResetCommand) Run(args []string) int {

	if len(args) < 2 {
		Info("You did not provided either branch or app name.")
		return 1
	}

	c.appName = args[0]
	var envBranch = args[1]
	if askForConfirmation("This will reset branch " + envBranch + " for " + c.appName + ". Continue?") {
		Info("Reseting branch %s for %s..", envBranch, c.appName)
		/*
			git branch -D  args[1]
			git checkout master upstream/master
			git push -f upstream/envBranch
		*/

		return 0
	} else {
		println("Reset aborted. No changes has been made.")
		return 0
	}
}

func (c *ResetCommand) Synopsis() string {
	return ""
}

func (c *ResetCommand) Help() string {
	helpText := `
Resets given environment to the latest master.
Note: it will remove history on branch and all merged branches.
`
	return strings.TrimSpace(helpText)
}
