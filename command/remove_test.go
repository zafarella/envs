package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestRemoveCommand_implement(t *testing.T) {
	var _ cli.Command = &RemoveCommand{}
}
