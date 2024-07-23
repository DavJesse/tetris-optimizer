package strung

import "testing"

func TestContains(t *testing.T) {
	subject := "..##"
	got := Contains(subject, "#")
	expected := true

	// Compare values of got and expected
	if got != true {
		t.Errorf("Got: %t", got)
		t.Errorf("Expected: %t", expected)
		t.Errorf("TestContains Failed!")
	}
}

func TestSplit(t *testing.T) {
	subject := "abcabcabc"
	found := Split(subject, "abc")
	expected := []string{"", "", "", ""}

	// Check if lengths of found and expected are equal...
	// ...or if their elements match
	for i := 0; i < len(expected); i++ {
		if len(found) != len(expected) || found[i] != expected[i] {
			// When lengths are different, print the following
			if len(found) != len(expected) {
				t.Errorf("Found has %d elements", len(found))
				t.Errorf("Expected has %d elements", len(expected))
				t.Errorf("TestSplitString Failed!")
				t.FailNow()

				// When elements don't match, print the following
			} else {
				t.Errorf("Found: %s", found)
				t.Errorf("Expected: %s", expected)
				t.Errorf("TestSplitString Failed!")

			}
		}
	}
}
