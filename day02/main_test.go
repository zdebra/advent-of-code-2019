package main

import (
	"fmt"
	"testing"
)

func Test_processIntcode(t *testing.T) {
	tests := []struct {
		inp []int
		pos0 int
	}{
		{
			inp:  []int{1, 0, 0, 0, 99},
			pos0: 2,
		},
		{
			inp:  []int{2,3,0,3,99},
			pos0: 2,
		},
		{
			inp:  []int{2,4,4,5,99,0},
			pos0: 2,
		},
		{
			inp:  []int{1,1,1,4,99,5,6,0,99},
			pos0: 30,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			processIntcode(tt.inp, 0)
			if tt.inp[0] != tt.pos0 {
				t.Errorf("expected pos0 to be %d was %d", tt.pos0, tt.inp[0])
				t.Log(tt.inp)
			}
		})
	}
}
