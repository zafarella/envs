package command

import (
	"strings"
)

type RemoveCommand struct {
	Meta
}

func (c *RemoveCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *RemoveCommand) Synopsis() string {
	return ""
}

func (c *RemoveCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
