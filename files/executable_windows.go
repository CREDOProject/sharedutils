package files

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// isExecutable returns true if the given file info is executable.
// On Windows, it just checks if the file extension is ".exe" or not.
func isExecutable(info fs.FileInfo) bool {
	return strings.ToLower(filepath.Ext(info.Name())) == ".exe"
}
