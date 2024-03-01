package files

import (
	"errors"
	"os"
	"path/filepath"
)

// execsInPath returns a list of executables in the given path.
// The returned paths are absolute.
//
// The given path should be an absolute path to a directory. If it's not
// a directory, the function will not proceed and return a nil slice.
//
// lookalike accepts a function that returns true if the file name matches what
// we are expecting.
// Example:
//
//	func looksLikePip(name string) bool {
//		var pipFileRegex = regexp.MustCompile(`^pip3(\d(\.\d\d?)?)?$`)
//		return pipFileRegex.MatchString(name)
//	}
func ExecsInPath(path string, lookalike func(string) bool) ([]string, error) {
	if !IsDir(path) {
		return nil, errors.New("Path is not a directory.")
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var execs []string

searchLoop:
	for _, entry := range entries {
		if entry.IsDir() || !lookalike(entry.Name()) {
			continue
		}

		resolvedPath, err := filepath.EvalSymlinks(filepath.Join(path, entry.Name()))
		if err != nil {
			return nil, err
		}
		info, err := os.Stat(resolvedPath)
		if err != nil {
			return nil, err
		}
		if !IsExecutable(info) {
			continue searchLoop
		}
		execs = append(execs, resolvedPath)
	}

	return execs, nil
}
