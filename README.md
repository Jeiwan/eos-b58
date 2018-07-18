base58 encoding/decoding for EOS
==========

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/btcsuite/btcutil/base58)

This package is fork of https://github.com/btcsuite/btcutil/tree/master/base58 which is adapted to the Base58 algorithm used by EOS (https://github.com/EOSIO/eos).

## Installation and Updating

```bash
$ go get -u github.com/Jeiwan/eos-b58
```

## Examples

* [Decode Example](https://github.com/Jeiwan/eos-b58/blob/master/example_test.go#L13)  
  Demonstrates how to decode modified base58 encoded data.
* [Encode Example](https://github.com/Jeiwan/eos-b58/blob/master/example_test.go#L26)  
  Demonstrates how to encode data using the modified base58 encoding scheme.
* [CheckDecode Example](https://github.com/Jeiwan/eos-b58/blob/master/example_test.go#L26)  
  Demonstrates how to decode Base58Check encoded data.
* [CheckEncode Example](https://github.com/Jeiwan/eos-b58/blob/master/example_test.go#L57)  
  Demonstrates how to encode data using the Base58Check encoding scheme.

## License

Package base58 is licensed under the [copyfree](http://copyfree.org) ISC
License.
