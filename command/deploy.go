package command

import (
	"fmt"
	"strings"
)

type DeployCommand struct {
	Meta
}

func (c *DeployCommand) Run(args []string) int {
	// Write your code here

	fmt.Println(args)
	//git merge args
	return 0
}

func (c *DeployCommand) Synopsis() string {
	return ""
}

func (c *DeployCommand) Help() string {
	helpText := `
 Deploys given branches to given environment/branch.
`
	return strings.TrimSpace(helpText)
}
