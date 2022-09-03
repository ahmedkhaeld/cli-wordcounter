package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {

	s := "word1 word2 word3 word4\n"
	b := bytes.NewBufferString(s)
	//words := 4
	//lines := 1
	bytes := 24
	res := count(b, false, true)
	if res != bytes {
		t.Errorf("Expected %d got %d instead.\n ", bytes, res)
	}
}
