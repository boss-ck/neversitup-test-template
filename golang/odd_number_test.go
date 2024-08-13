package main

import "testing"

// Write your test here

func TestFindOddNumber(t *testing.T) {
	type args struct {
		text []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1 - [7]",
			args: args{
				text: []int{7},
			},
			want: 7,
		},
		{
			name: "case 2 - [0]",
			args: args{
				text: []int{0},
			},
			want: 0,
		},
		{
			name: "case 3 - [1,1,2]",
			args: args{
				text: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "case 4 - [0,1,0,1,0]",
			args: args{
				text: []int{0, 1, 0, 1, 0},
			},
			want: 0,
		},
		{
			name: "case 5 - [1,2,2,3,3,3,4,3,3,3,2,2,1]",
			args: args{
				text: []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			},
			want: 4,
		},
		{
			name: "case 6 - [1,1,2,2] เคสที่ไม่เจอที่เป็นเลขคี่เลยให้ return -1 แทน",
			args: args{
				text: []int{1, 1, 2, 2},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindOddNumber(tt.args.text); got != tt.want {
				t.Errorf("FindOddNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
