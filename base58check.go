// Copyright (c) 2013-2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package base58

import (
	"bytes"
	"errors"

	"golang.org/x/crypto/ripemd160"
)

// ErrChecksum indicates that the checksum of a check-encoded string does not verify against
// the checksum.
var ErrChecksum = errors.New("checksum error")

// ErrInvalidFormat indicates that the check-encoded string has an invalid format.
var ErrInvalidFormat = errors.New("invalid format: version and/or checksum bytes missing")

var ErrUnsupportedType = errors.New("only K1 type is supported")

// checksum: first four bytes of sha256^2
func checksum(input []byte) (cksum [4]byte) {
	ripe := ripemd160.New()
	ripe.Write(input)
	h := ripe.Sum(nil)
	copy(cksum[:], h[:4])
	return
}

// CheckEncode prepends a version byte and appends a four byte checksum.
func CheckEncode(input []byte) string {
	b := make([]byte, 0, len(input)+4)
	b = append(b, input[:]...)
	cksum := checksum(b)
	b = append(b, cksum[:]...)
	return Encode(b)
}

// CheckDecode decodes a string that was encoded with CheckEncode and verifies the checksum.
func CheckDecode(input string) (result []byte, err error) {
	decoded := Decode(input)
	if len(decoded) < 4 {
		return nil, ErrInvalidFormat
	}
	var cksum [4]byte
	copy(cksum[:], decoded[len(decoded)-4:])
	if checksum(decoded[:len(decoded)-4]) != cksum {
		return nil, ErrChecksum
	}
	payload := decoded[:len(decoded)-4]
	result = append(result, payload...)
	return
}

// CheckDecodeWithType decodes a string that was encoded with CheckEncode and verifies the checksum.
func CheckDecodeWithType(input, typ string) (result []byte, err error) {
	if typ != "K1" {
		return nil, ErrUnsupportedType
	}

	decoded := Decode(input)
	proofChecksum := decoded[len(decoded)-4:]
	data := decoded[:len(decoded)-4]

	checkData := [][]byte{data}
	checkData = append(checkData, []byte(typ))

	actualChecksum := checksum(bytes.Join(checkData, []byte{}))
	if bytes.Compare(actualChecksum[:], proofChecksum) != 0 {
		return nil, ErrChecksum
	}

	payload := decoded[:len(decoded)-4]
	result = append(result, payload...)

	return
}
