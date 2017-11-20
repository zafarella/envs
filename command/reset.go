package command

import (
	"strings"
)

type ResetCommand struct {
	Meta
}

func (c *ResetCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *ResetCommand) Synopsis() string {
	return ""
}

func (c *ResetCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
