package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestAllocSignalCommand_Implements(t *testing.T) {
	t.Parallel()
	var _ cli.Command = &AllocSignalCommand{}
}
