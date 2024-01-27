package code

import "testing"

func TestGenerateAllCodes(t *testing.T) {
	expectedSize := 6720
	expectedFirstCode := Code{0, 0, 0, 0, 0}
	expectedLastCode := Code{7, 7, 7, 7, 7}
	actualCodes := GenerateAllCodes()
	actualSize := len(actualCodes)

	if actualSize != expectedSize {
		t.Fatalf("expected slice with a size of %v, but was size of %v", expectedSize, actualSize)
	}
	if !codesAreEqual(actualCodes[0], expectedFirstCode) {
		t.Fatalf("expected first code %v, got %v", expectedFirstCode, actualCodes[0])
	}
	if !codesAreEqual(actualCodes[6719], expectedLastCode) {
		t.Fatalf("expected last code %v, got %v", expectedLastCode, actualCodes[6719])
	}

}

func codesAreEqual(code1 Code, code2 Code) bool {
	if len(code1) != len(code2) {
		return false
	}

	for i := 0; i < len(code1); i++ {
		if code1[i] != code2[i] {
			return false
		}
	}

	return true
}
