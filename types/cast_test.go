package types

import "testing"

type Hello struct {
	Name string
}

func Test_To(t *testing.T) {
	hello := Hello{Name: "World"}

	converted, err := To[Hello](hello)
	if err != nil {
		t.Error(err)
	}
	if converted.Name != "World" {
		t.Error("Expected World.")
	}
}
