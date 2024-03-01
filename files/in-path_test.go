package files

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
	"testing"
)

var testRegex = regexp.MustCompile(`^test(\d(\.\d\d?)?)?$`)

func looksLike(name string) bool {
	return testRegex.MatchString(name)
}

func TestExecsInPath(t *testing.T) {
	tmpDir := t.TempDir()

	for _, file := range []string{"test0", "test1", "notacandidate"} {
		filePath := filepath.Join(tmpDir, file)
		_, err := os.Create(filePath)
		if err != nil {
			t.Errorf("Got error creating file: %s", file)
		}
		err = os.Chmod(filePath, 0700)
		if err != nil {
			t.Errorf("Got error changing permission to file: %s", file)
		}
		defer os.Remove(filePath)
	}

	test2 := "test2"
	test2path := path.Join(tmpDir, test2)
	os.Create(test2path)
	defer os.Remove(test2path)

	unreadablepath := path.Join(tmpDir, "unreadable")
	os.Mkdir(unreadablepath, 0200)

	foundExecutables, err := ExecsInPath(tmpDir, looksLike)
	if err != nil {
		t.Errorf("ExecsInPath(%s) returned error: %v, want nil", tmpDir, err)
	}
	expectedExecutables := 2
	if len(foundExecutables) != expectedExecutables {
		t.Errorf("ExecsInPath(%s) returned %d executables, want %d", tmpDir, len(foundExecutables), expectedExecutables)
	}
	// Testing without a dir
	_, err = ExecsInPath(test2path, looksLike)
	if err == nil {
		t.Errorf("ExecsInPath(%s) has't returned error.", test2path)
	}
	// Testing unreadable path
	_, err = ExecsInPath(unreadablepath, looksLike)
	if err == nil {
		t.Errorf("ExecsInPath(%s) has't returned error.", test2path)
	}
	test3path := path.Join(tmpDir, "test3")
	os.Symlink("buh", test3path)

	_, err = ExecsInPath(tmpDir, looksLike)
	if err == nil {
		t.Errorf("ExecsInPath(%s) has't returned error.", test2path)
	}
}
