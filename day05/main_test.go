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
