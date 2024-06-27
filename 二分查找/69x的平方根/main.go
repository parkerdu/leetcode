package main

func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		mid := (l + r) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
