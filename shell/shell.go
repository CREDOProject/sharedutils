package shell

import "os/exec"

// ExecShim provides an abstraction layer to use the os/exec API.
//
// Implementation of ExecShim MUST use os/exec (or an equivalent) if they act
// on the host system.
type ExecShim interface {
	LookPath(string) (string, error)
	Command(string, ...string) *exec.Cmd
}

// A execShellCommand represent a shell comand to be acted upon.
// It provides the implementation of an ExecShim.
type execShellCommand struct{}

// LookPath searches for an executable named file in the
// directories named by the PATH environment variable.
// If file contains a slash, it is tried directly and the PATH is not consulted.
// Otherwise, on success, the result is an absolute path.
func (exc execShellCommand) LookPath(name string) (string, error) {
	return exec.LookPath(name)
}

// Command returns the Cmd struct to execute the named program with
// the given arguments.
func (exc execShellCommand) Command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}

// Returns a new execShellCommand.
func New() ExecShim {
	return execShellCommand{}
}
