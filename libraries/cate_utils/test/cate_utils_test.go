package cateutils

import (
	"testing"
)

func TestCateUtils(t *testing.T) {
	result := CateUtils("works")
	if result != "CateUtils works" {
		t.Error("Expected CateUtils to append 'works'")
	}
}
