package main

import "testing"

func TestDouble(t *testing.T) {
	var tests = []struct {
		name         string
		input        int
		expectOutput int
	}{
		{
			name:         "double 1 should got 2",
			input:        1,
			expectOutput: 2,
		},
		{
			name:         "double 2 should got 4",
			input:        2,
			expectOutput: 4,
		},
	}

	for _, test := range tests {
		tc:=test // 显式绑定，每次循环都会生成一个新的tc变量
		t.Run(test.name, func(t *testing.T) {
			if test.expectOutput != Double(tc.input) {
				t.Fatalf("expect: %d, but got: %d", test.input, test.expectOutput)
					}
		})
	}
}