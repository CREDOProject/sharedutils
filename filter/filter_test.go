package filter

import "testing"

func always(a string) bool { return true }
func onlya(a string) bool  { return a == "a" }

func Test_Filter(t *testing.T) {
	a := []string{"a", "b", "c"}
	filtered := Filter(a, always, onlya)
	if len(filtered) != 1 {
		t.Error("Oops")
	}
}
