package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCut(t *testing.T) {
	tests := []struct {
		name        string
		flags       Flags
		input, want []string
	}{
		{
			name:  "cut with multiple fields",
			flags: Flags{f: "1,3", s: true, d: "\t"},
			input: []string{"aaa\tccc\tbbb"},
			want:  []string{"aaa\tbbb"},
		},
		{
			name:  "cut with different delimiter",
			flags: Flags{f: "1", s: true, d: ":"},
			input: []string{"Spring:green:grass:warm"},
			want:  []string{"Spring"},
		},
		{
			name:  "cut with different range",
			flags: Flags{f: "1-3", s: true, d: ":"},
			input: []string{"Spring:green:grass:warm"},
			want:  []string{"Spring:green:grass"},
		},
		{
			name:  "cut with multiple fields",
			flags: Flags{f: "1,3,4", s: true, d: ":"},
			input: []string{"Spring:green:grass:warm"},
			want:  []string{"Spring:grass:warm"},
		},
		{
			name:  "cut with fields beyond column length",
			flags: Flags{f: "6,7", s: true, d: ":"},
			input: []string{"Spring:green:grass:warm"},
			want:  []string{""},
		},
		{
			name:  "cut with fields beyond column length",
			flags: Flags{f: "8-9", s: true, d: ":"},
			input: []string{"Spring:green:grass:warm"},
			want:  []string{""},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := cut(tc.input, tc.flags)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}
