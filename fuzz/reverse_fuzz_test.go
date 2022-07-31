/**
The unit test has limitations, namely that each input must be added to the test by the developer.
One benefit of fuzzing is that it comes up with inputs for your code, and may identify edge cases that the test cases you came up with didn’t reach.
**/
package main

import (
	"testing"
	"unicode/utf8"
)

// go test -run=FuzzReverse
// go test -run=FuzzReverse -fuzz=Fuzz  // run with fuzz, The default is to run forever if no failures occur, can be interrupted with ctrl-C.
// go test -run=FuzzReverse -fuzz=Fuzz -fuzztime 30s  // fuzz for 30 seconds before exiting if no failure was found.
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
