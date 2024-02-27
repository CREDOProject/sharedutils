package files

import (
	"os"
	"time"
)

type mockFileInfo struct {
	name      string
	mode      os.FileMode
	isDir     bool
	isRegular bool
}

func (m mockFileInfo) Name() string {
	return m.name
}

func (m mockFileInfo) Size() int64 {
	return 0
}

func (m mockFileInfo) Mode() os.FileMode {
	return m.mode
}

func (m mockFileInfo) ModTime() time.Time {
	return time.Time{}
}

func (m mockFileInfo) IsDir() bool {
	return m.isDir
}

func (m mockFileInfo) Sys() interface{} {
	return nil
}
