package main

import (
	"testing"
)

func TestHello2(tis *testing.T) {
	res := GetHello2()
	// require.Equal(tis, "hello world 1", res)
	if res != "hello world 1" {
		tis.Fail()
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		//TODO : Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
