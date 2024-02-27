package files

import (
	"os"
	"runtime"
	"testing"
)

func Test_Executable(t *testing.T) {
	tests := []struct {
		fileInfo mockFileInfo
		expected bool
	}{}

	switch runtime.GOOS {
	default:
		tests = []struct {
			fileInfo mockFileInfo
			expected bool
		}{
			// Test for regular executable file on Unix-like systems
			{mockFileInfo{"test", 0755, false, true}, true},
			// Test for regular non-executable file on Unix-like systems
			{mockFileInfo{"test", 0644, false, true}, false},
			// Test for directory
			{mockFileInfo{"test", 0755 | os.ModeDir, true, false}, false},
			// Test for regular executable file on Windows
			//,
			// Test for regular non-executable file on Windows
			{mockFileInfo{"test.txt", 0, false, true}, false},
		}
	case "windows":
		tests = []struct {
			fileInfo mockFileInfo
			expected bool
		}{
			{mockFileInfo{"test.exe", 0, false, true}, true},
			{mockFileInfo{"test.bat", 0, false, true}, false},
		}
	}

	for _, test := range tests {
		result := IsExecutable(test.fileInfo)
		if result != test.expected {
			t.Errorf(
				"For file %s, expected %t but got %t",
				test.fileInfo.Name(), test.expected, result)
		}
	}

}
