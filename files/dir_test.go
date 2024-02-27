package files

import (
	"os"
	"testing"
)

func Test_IsDir(t *testing.T) {
	temp := os.TempDir()

	isDir := IsDir(temp)
	if !isDir {
		t.Errorf(
			"For directory %s, expected %t but got %t",
			temp, true, false)
	}
}
