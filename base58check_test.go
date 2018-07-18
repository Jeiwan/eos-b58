// Copyright (c) 2013-2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package base58_test

import (
	"testing"

	"github.com/Jeiwan/eos-b58"
)

var checkEncodingStringTests = []struct {
	in  string
	out string
}{
	{"", "4zNxKW"},
	{" ", "4goVXTs"},
	{"-", "67QwLzu"},
	{"0", "6VrpPNL"},
	{"1", "6cfKu5c"},
	{"-1", "PWLSUnKv"},
	{"11", "RVqS4VrT"},
	{"abc", "4h3c82uKAN"},
	{"1234598760", "K5zqBMZZTzUbAbtDAhc"},
	{"abcdefghijklmnopqrstuvwxyz", "LWmP1W82eUos2HWzVn19rapmig4X5dqPWgGGVdnb9"},
	{"00000000000000000000000000000000000000000000000000000000000000", "KmcpTepcsgVkLNE6BaRjwbfCHAq2dBLFiXZcTrsZGjmGXTgquAWguHxbxzF8zuiobsc8ix9VpBmrjsQ4KUfoAFZEmB"},
}

func TestBase58Check(t *testing.T) {
	for x, test := range checkEncodingStringTests {
		// test encoding
		if res := base58.CheckEncode([]byte(test.in)); res != test.out {
			t.Errorf("CheckEncode test #%d failed: got %s, want: %s", x, res, test.out)
		}

		// test decoding
		res, err := base58.CheckDecode(test.out)
		if err != nil {
			t.Errorf("CheckDecode test #%d failed with err: %v", x, err)
		} else if string(res) != test.in {
			t.Errorf("CheckDecode test #%d failed: got: %s want: %s", x, res, test.in)
		}
	}

	// test the two decoding failure cases
	// case 1: checksum error
	_, err := base58.CheckDecode("3MNQE1Y")
	if err != base58.ErrChecksum {
		t.Error("Checkdecode test failed, expected ErrChecksum")
	}
	// case 2: invalid formats (string lengths below 5 mean the version byte and/or the checksum
	// bytes are missing).
	testString := ""
	for len := 0; len < 4; len++ {
		// make a string of length `len`
		_, err = base58.CheckDecode(testString)
		if err != base58.ErrInvalidFormat {
			t.Error("Checkdecode test failed, expected ErrInvalidFormat")
		}
	}

}
