package main

// /*
// Binary search => need a hi, lo, midpoint
// check if mid equals target => evaluate hi and lo to v to adjust midpoint.
// 	when adjusting midpoint, make sure not to include again
// 	take: [lo, hi)
// */

// func BinarySearch(xi []int, t int) int {
// 	hi := len(xi)
// 	lo := 0

// 	m := (lo + (hi - lo)) / 2
// 	v := xi[m]

// 	for lo < hi {
// 		if v == t {
// 			return m
// 		} else if v > t {
// 			hi = m // means that the target will be in the lower half (hi = m because hi is exclusive)
// 		} else {
// 			lo = m + 1 //means that the mid is less than our target. Our target is in the larger half. Add 1 to not have to check the mid again [lo is inclusive]
// 		}
// 	}
// 	return -1
// }
func BinarySearch(xi []int, t int) int {
	hi := len(xi)
	lo := 0

	for lo < hi {
		m := lo + (hi-lo)/2
		v := xi[m]
		if v == t {
			return m
		} else if v > t {
			hi = m // means that the target will be in the lower half (hi = m because hi is exclusive)
		} else {
			lo = m + 1 //means that the mid is less than our target. Our target is in the larger half. Add 1 to not have to check the mid again [lo is inclusive]
		}
	}
	return -1
}
