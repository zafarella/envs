package command

import (
	"strings"
)

type UpdateCommand struct {
	Meta
}

func (c *UpdateCommand) Run(args []string) int {
	/*
		1.compare hashes/tags
		2. if diff
		3. download new version
	*/
	return 0
}

func (c *UpdateCommand) Synopsis() string {
	return ""
}

func (c *UpdateCommand) Help() string {
	helpText := `
Updates the envs binary to latest available version.

`
	return strings.TrimSpace(helpText)
}
