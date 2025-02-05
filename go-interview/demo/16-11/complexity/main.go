package main

// T1(n) = O(n)，T2(n) = O(n^2 )，则 T1(n) +T2(n) = O(n^2 )
// 复杂度 T(n)=100000000+n+n
// 最终复杂度： T(n)=O（n）
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

// T1(n) = O(n)，T2(n) = O(n^2)，则 T1(n) * T2(n) = O(n^3)
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
