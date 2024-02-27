package shell

import (
	"testing"
)

func Test_New(t *testing.T) {
	exec := New()

	if exec == nil {
		t.Errorf("Not returning execShellCommand")
	}
}

func Test_LookPath(t *testing.T) {
	exec := New()

	_, err := exec.LookPath("ls")
	if err != nil {
		t.Errorf("LookPath failed with error: %v", err)
	}
}

func Test_Exec(t *testing.T) {
	exec := New()

	cmd := exec.Command("ls", "-l")
	if cmd == nil {
		t.Error("Command returned nil")
	}

	if len(cmd.Args) != 2 || cmd.Args[1] != "-l" {
		t.Errorf("Command arguments expected ['-l'], got %v", cmd.Args)
	}
}
