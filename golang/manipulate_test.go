package main

import (
	"reflect"
	"testing"
)

// Write your test here
func TestManipulate(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "case 1 - a",
			args: args{
				text: "a",
			},
			want: []string{"a"},
		},
		{
			name: "case 2 - ab",
			args: args{
				text: "ab",
			},
			want: []string{"ab", "ba"},
		},
		{
			name: "case 3 - abc",
			args: args{
				text: "abc",
			},
			want: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			name: "case 4 - abab",
			args: args{
				text: "abab",
			},
			want: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Manipulate(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manipulate() = %v, want %v", got, tt.want)
			}
		})
	}
}
