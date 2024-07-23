package strung

import "testing"

func TestContains(t *testing.T) {
	subject := "..##"
	got := Contains(subject, "#")
	expected := true

	if got != true {
		t.Errorf("Got: %t", got)
		t.Errorf("Expected: %t", expected)
		t.Errorf("TestContains Failed!")
	}
}
