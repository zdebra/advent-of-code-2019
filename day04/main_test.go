package main

import "testing"

func Test_matchCriteria(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "111111",
			n:    111111,
			want: true,
		},
		{
			name: "123456",
			n:    123456,
			want: false,
		},
		{
			name: "122256",
			n:    122256,
			want: true,
		},
		{
			name: "122222",
			n:    122222,
			want: true,
		},
		{
			name: "122252",
			n:    122252,
			want: false,
		},
		{
			name: "223450",
			n:    223450,
			want: false,
		},
		{
			name: "123789",
			n:    123789,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchCriteria(tt.n); got != tt.want {
				t.Errorf("matchCriteria() = %v, want %v", got, tt.want)
			}
		})
	}
}
