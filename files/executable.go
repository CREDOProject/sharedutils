//go:build unix

package files

import (
	"io/fs"
)

// IsExecutable returns true if the given file info is executable.
// On Windows, it just checks if the file extension is ".exe" or not.
func IsExecutable(info fs.FileInfo) bool {
	return info.Mode().IsRegular() && info.Mode()&0o111 != 0
}
