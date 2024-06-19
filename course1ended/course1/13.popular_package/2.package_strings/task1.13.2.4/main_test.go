package main

import (
	"testing"
	"unicode/utf8"
)

func TestGenerateActivationKey(t *testing.T) {
	got := generateActivationKey()
	if 19 != utf8.RuneCountInString(got) {
		t.Errorf("expected = %d, got = %d", 19, len(got))
	}
}
