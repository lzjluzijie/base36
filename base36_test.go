package base36_test

import (
	"bytes"
	"testing"

	"github.com/lzjluzijie/base36"
)

var tests = []struct {
	data    []byte
	encoded string
}{
	{[]byte(""), ""},
	{[]byte(" "), "w"},
	{[]byte("-"), "19"},
	{[]byte("0"), "1c"},
	{[]byte("-1"), "8xd"},
	{[]byte("11"), "9pt"},
	{[]byte("abc"), "3ssir"},
	{[]byte("1234567890"), "11txtitqhl6b2o1c"},
	{[]byte("Hello World!"), "2678lx5gvmsv1dro9b5"},
	{[]byte("https://halu.lu/"), "66mfwrmo1x1vcu4q34y5nufjj"},
	{[]byte("abcdefghijklmnopqrstuvwxyz"), "vizqx2p66dy1mjoqmz90ven2q0rrb12x1ikq44fu"},
	{[]byte("00000000000000000000000000000000000000000000000000000000000000"), "5gefrzx95jt742iv4o7fgpj4dnzn1162jv8eiaqqfg7sz5cm07qyhlhrdvldhnni6esrlolznw6zu6899k0j36eq2klveqeo"},
}

var invalidTests = []string{
	" ",
	"A",
	"-",
	"+",
	"=",
	"1+1>2",
	"233 333",
	"哈陆lu",
	"https://github.com/lzjluzijie/base36",
}

func TestBase58(t *testing.T) {
	// Encode tests
	for x, test := range tests {
		if res := base36.Encode(test.data); res != test.encoded {
			t.Errorf("Encode test #%d failed: got: %s want: %s",
				x, res, test.encoded)
			continue
		}
	}

	// Decode tests
	for x, test := range tests {
		if res := base36.Decode(test.encoded); !bytes.Equal(res, test.data) {
			t.Errorf("Decode test #%d failed: got: %q want: %q",
				x, test.encoded, res)
			continue
		}
	}

	// Decode with invalid input
	for x, test := range invalidTests {
		if res := base36.Decode(test); res != nil {
			t.Errorf("Decode invalidString test #%d failed: %q",
				x, res)
			continue
		}
	}
}
