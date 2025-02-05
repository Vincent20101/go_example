package main

import (
	"testing"
	"time"
)

func TestDouble(t *testing.T) {
	//t.Parallel()

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

	for k, test := range tests {
		// 如果使用子测试（使用 t.Run），每个子测试也可以调用 t.Parallel() 来并行执行
		//If use t.Parallel(), best to use the following code
		tc := test
		kk := k
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if kk == 0 {
				time.Sleep(5 * time.Second)
			} else {
				time.Sleep(2 * time.Second)
			}
			if tc.expectOutput != Double(tc.input) {
				t.Fatalf("expect: %d, but got: %d", tc.input, tc.expectOutput)
			}
		})

		//t.Run(test.name, func(t *testing.T) {
		//	if test.expectOutput != Double(test.input) {
		//		t.Fatalf("expect: %d, but got: %d", test.input, test.expectOutput)
		//	}
		//})
	}
}

// 在这个示例中，三个测试函数都调用了 t.Parallel()，因此它们会并行执行。由于每个测试都有不同的休眠时间，
// 总的测试时间将取决于最长的测试（3 秒），而不是所有测试时间的总和（6 秒）。
func TestAddition1(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second) // 模拟耗时操作
	if 1+1 != 2 {
		t.Error("1+1 should be 2")
	}
}

func TestAddition2(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second) // 模拟耗时操作
	if 2+2 != 4 {
		t.Error("2+2 should be 4")
	}
}

func TestAddition3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second) // 模拟耗时操作
	if 3+3 != 6 {
		t.Error("3+3 should be 6")
	}
}
