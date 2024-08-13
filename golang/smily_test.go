package main

import "testing"

// Write your test here

func TestCountSmilyFace(t *testing.T) {

	// countSmileys([':)', ';(', ';}', ':-D']);       // should return 2;
	// countSmileys([';D', ':-(', ':-)', ';~)']);     // should return 3;
	// countSmileys([';]', ':[', ';*', ':$', ';-D']); // should return 1;

	type args struct {
		text []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1 - [':)', ';(', ';}', ':-D']",
			args: args{
				text: []string{":)", ";(", ";}", ":-D"},
			},
			want: 2,
		},
		{
			name: "case 2 - [';D', ':-(', ':-)', ';~)']",
			args: args{
				text: []string{";D", ":-(", ":-)", ";~)"},
			},
			want: 3,
		},
		{
			name: "case 3 - [';]', ':[', ';*', ':$', ';-D']",
			args: args{
				text: []string{";]", ":[", ";*", ":$", ";-D"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSmilyFace(tt.args.text); got != tt.want {
				t.Errorf("CountSmilyFace() = %v, want %v", got, tt.want)
			}
		})
	}
}
