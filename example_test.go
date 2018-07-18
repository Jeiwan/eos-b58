// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package base58_test

import (
	"fmt"

	"github.com/Jeiwan/eos-b58"
)

// This example demonstrates how to decode modified base58 encoded data.
func ExampleDecode() {
	// Decode example modified base58 encoded data.
	encoded := "25JnwSn7XKfNQ"
	decoded := base58.Decode(encoded)

	// Show the decoded data.
	fmt.Println("Decoded Data:", string(decoded))

	// Output:
	// Decoded Data: Test data
}

// This example demonstrates how to encode data using the modified base58
// encoding scheme.
func ExampleEncode() {
	// Encode example data with the modified base58 encoding scheme.
	data := []byte("Test data")
	encoded := base58.Encode(data)

	// Show the encoded data.
	fmt.Println("Encoded Data:", encoded)

	// Output:
	// Encoded Data: 25JnwSn7XKfNQ
}

// This example demonstrates how to decode Base58Check encoded data.
func ExampleCheckDecode() {
	// Decode an example Base58Check encoded data.
	encoded := "A1zP1eP5QGefi2DMPTfTL5SLmv7EpHB2E"
	decoded, err := base58.CheckDecode(encoded)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Show the decoded data.
	fmt.Printf("Decoded data: %x\n", decoded)

	// Output:
	// Decoded data: 62e907b15cbf27d5425399ebf6f0fb50ebb88f18
}

// This example demonstrates how to encode data using the Base58Check encoding
// scheme.
func ExampleCheckEncode() {
	// Encode example data with the Base58Check encoding scheme.
	data := []byte("Test data")
	encoded := base58.CheckEncode(data)

	// Show the encoded data.
	fmt.Println("Encoded Data:", encoded)

	// Output:
	// Encoded Data: 82iP79GRURMpCWQeS2
}
