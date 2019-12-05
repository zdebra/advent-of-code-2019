package main

import (
	"reflect"
	"testing"
)

func Test_decodeInstruction(t *testing.T) {
	type args struct {
		code int
	}
	tests := []struct {
		name  string
		code  int
		want  instruction
		want1 [3]int
	}{
		{
			code:  1002,
			want:  multiply,
			want1: [3]int{0, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := decodeInstruction(tt.code)
			if got != tt.want {
				t.Errorf("decodeInstruction() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("decodeInstruction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_program_run(t *testing.T) {

	tests := []struct {
		name string
		p    *program
		want int
	}{
		{
			name: "less than 8 A",
			p: &program{
				memory: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
				input:  3,
			},
			want: 1,
		},
		{
			name: "less than 8 B",
			p: &program{
				memory: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
				input:  9,
			},
			want: 0,
		},
		{
			name: "larger example A",
			p: &program{
				memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
					1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
					999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
				input: 7,
			},
			want: 999,
		},
		{
			name: "larger example B",
			p: &program{
				memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
					1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
					999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
				input: 8,
			},
			want: 1000,
		},
		{
			name: "larger example C",
			p: &program{
				memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
					1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
					999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
				input: 10,
			},
			want: 1001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.p.run(0)
			if tt.p.output != tt.want {
				t.Errorf("got %d expected %d", tt.p.output, tt.want)
			}
		})
	}
}
