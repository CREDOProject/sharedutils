package files

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

var testRegex = regexp.MustCompile(`^test(\d(\.\d\d?)?)?$`)

func looksLike(name string) bool {
	return testRegex.MatchString(name)
}

func TestExecsInPath(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Create some test files with different extensions
	executableFiles := []string{"test1", "test2", "test3.exe", "test4", "test5.sh"}
	for _, file := range executableFiles {
		filePath := filepath.Join(tmpDir, file)
		os.Create(filePath)
		if filepath.Ext(file) != ".exe" {
			os.Chmod(filePath, 0755)
		}
		defer os.Remove(filePath)
	}

	// Test when the directory doesn't exist
	nonExistentDir := filepath.Join(tmpDir, "non-existent-dir")
	res, err := ExecsInPath(nonExistentDir, looksLike)
	if err != nil && res != nil {
		t.Errorf("ExecsInPath(%s) returned error: %v, want nil", nonExistentDir, err)
	}

	// Test when the directory is empty
	emptyDir := t.TempDir()
	execs, err := ExecsInPath(emptyDir, looksLike)
	if err != nil {
		t.Errorf("ExecsInPath(%s) returned error: %v, want nil", emptyDir, err)
	}
	if len(execs) != 0 {
		t.Errorf("ExecsInPath(%s) returned %d executables, want 0", emptyDir, len(execs))
	}

	// Test when the directory contains executable files
	foundExecutables, err := ExecsInPath(tmpDir, looksLike)
	if err != nil {
		t.Errorf("ExecsInPath(%s) returned error: %v, want nil", tmpDir, err)
	}
	expectedExecutables := 3
	if len(foundExecutables) != expectedExecutables {
		t.Errorf("ExecsInPath(%s) returned %d executables, want %d", tmpDir, len(foundExecutables), expectedExecutables)
	}
}
