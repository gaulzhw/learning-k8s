package mutex

import (
	"testing"
)

func TestSafeCount(t *testing.T) {
	expected := uint64(1000000)
	got := safeCount()
	if got != expected {
		t.Errorf("safe count should %d, but got %d", expected, got)
	}
}
