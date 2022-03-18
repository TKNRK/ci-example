package main

import (
	"fmt"
	"testing"
)

func TestInclude(t *testing.T) {
	type input struct {
		slice  []int
		target int
	}
	type test struct {
		desc  string
		input input
		want  bool
	}

	tests := []test{
		{
			desc: "empty slice includes nothing",
			input: input{
				slice:  []int{},
				target: 1,
			},
			want: false,
		},
		{
			desc: "slice includes target value",
			input: input{
				slice:  []int{-1, 0, 1},
				target: 1,
			},
			want: true,
		},
		{
			desc: "slice does not include target value",
			input: input{
				slice:  []int{1, 2, 3},
				target: 99,
			},
			want: false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", i, tt.desc), func(t *testing.T) {
			got := include(tt.input.slice, tt.input.target)
			if got != tt.want {
				t.Errorf("want: %v, but got: %v", tt.want, got)
			}
		})
	}
}
