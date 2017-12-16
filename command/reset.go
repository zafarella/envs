package command

import (
	"github.com/fatih/color"
	"strings"
)

type ResetCommand struct {
	Meta
}

func (c *ResetCommand) Run(args []string) int {
	red := color.New(color.FgHiRed).PrintfFunc()

	if len(args) < 2 {
		red("You missing either branch or app name.")
		return 1
	}

	var app string = args[0]
	var envBranch string = args[1]
	if askForConfirmation("This will reset branch " + envBranch + " for " + app + ". Continue?") {
		red("Reseting branch %s for %s..", envBranch, app)
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
