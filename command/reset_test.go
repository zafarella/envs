package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestResetCommand_implement(t *testing.T) {
	var _ cli.Command = &ResetCommand{}
}
