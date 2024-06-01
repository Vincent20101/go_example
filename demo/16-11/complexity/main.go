package main

// 复杂度 T(n)=100000000+n+n
// 最终复杂度： T(n)=n
func f1(n int) {
	for i := 0; i < 100000000; i++ {
		// 执行一些操作
	}
	for i := 0; i < n; i++ {
		// 执行一些操作
	}
	for i := 0; i < n; i++ {
		// 执行一些操作
	}
}

// 复杂度 T(n)=100000000+n*n
// 最终复杂度： T(n)=n^2
func f2(n int) {
	for i := 0; i < 100000000; i++ {
		// 执行一些操作
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; i++ {
			// 执行一些操作
		}
	}
}

func f3(n int) {
	s := []int{}
	x := 10
	for i := 0; i < n; i++ {
		s[i] = i * x
	}
}
