package testhelper

import (
	"testing"
)

func TestTestHelper(t *testing.T) {
	result := TestHelper("works")
	if result != "TestHelper works" {
		t.Error("Expected TestHelper to append 'works'")
	}
}
