package command

import (
	"bytes"
	"fmt"
)

type VersionCommand struct {
	Meta

	Name      string
	Revision  string
	Version   string
	BuildBy string
	Tag    string
}

func (c *VersionCommand) Run(args []string) int {
	var versionString bytes.Buffer

	fmt.Fprintf(&versionString, "%s version ", c.Name, c.Version)
	if c.Revision != "" {
		fmt.Fprintf(&versionString, " Revision : %s", c.Revision)
	}
	if c.BuildBy != "" {
		fmt.Fprintf(&versionString, " Build by : %s", c.BuildBy)
	}
	if c.Tag != "" {
		fmt.Fprintf(&versionString, " Tag : %s", c.Tag)
	}

	c.Ui.Output(versionString.String())
	return 0
}

func (c *VersionCommand) Synopsis() string {
	return fmt.Sprintf("Print %s version and quit", c.Name)
}
func (c *VersionCommand) Help() string {
	return ""
}
