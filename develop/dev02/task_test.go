package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	type test struct {
		input, want string
	}
	tests := []test{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
		{`qwe\4\5`, "qwe45"},
		{`qwe\45`, "qwe44444"},
		{`qwe\\5`, `qwe\\\\\`},
		{`45fdvvf234`, `fdvvff`},
		{`4f5dvvf234`, `fffffdvvff`},
		{`abc123sds`, `abcsds`},
		{`abc023sds`, `absds`},
	}

	for _, tc := range tests {
		got, _ := UnpackString(tc.input)
		if tc.want != got {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
