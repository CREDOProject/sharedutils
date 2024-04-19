package files

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

// FilesInPath returns a list of files in the given path.
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
//
// additionalChecks accepts a list of functions that returns true if the file
// meets the requirements.
// Example:
//
//	func IsExecutable(info fs.FileInfo) bool {
//		return strings.ToLower(filepath.Ext(info.Name())) == ".exe"
//	}
func FilesInPath(path string,
	lookalike func(string) bool,
	additionalChecks ...func(fs.FileInfo) bool) ([]string, error) {
	if !IsDir(path) {
		return nil, errors.New("Path is not a directory.")
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var execs []string

outerloop:
	for _, entry := range entries {
		if entry.IsDir() || !lookalike(entry.Name()) {
			continue
		}
		entryPath := filepath.Join(path, entry.Name())
		resolvedPath, err := filepath.EvalSymlinks(entryPath)
		if err != nil {
			return nil, err
		}
		info, err := os.Stat(resolvedPath)
		if err != nil {
			return nil, err
		}
		for _, check := range additionalChecks {
			if !check(info) {
				continue outerloop
			}

		}
		execs = append(execs, resolvedPath)
	}
	return execs, nil
}

// ExecsInPath returns a list of executables in the given path.
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
	return FilesInPath(path, lookalike, IsExecutable)
}
