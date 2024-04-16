package main

import (
	"math"
)

//given two crystal balls will break from high enough distance, determine the exact spot in which it will break in the most optimized way.
//O(sqrt(N)) => jump forward until both break. then walk back to last jump
// func Break(o []bool) int {
// 	const oLen int = oLen
// 	jmpLength := int(math.Sqrt(float64(oLen)))

// 	for i := jmpLength; i < oLen; i += jmpLength {
// 		if o[i] {
// 			break
// 		}
// 	}
// 	i -= jmpAmount;

// 	for j:=0; j < jmpAmount && i < oLen; j++, i++{
// 		if o[i] {
// 			return i
// 		}
// 	}
// 	return -1
// }

// Break finds the exact spot where two crystal balls break
// from a high enough distance in the most optimized way.
// It returns the spot where the balls break.
// Complexity: O(sqrt(N)) => jump forward until both break. then walk back to last jump.
func Break(o []bool) int {
	oLen := len(o)
	jmpLength := int(math.Sqrt(float64(oLen)))

	// Jump forward until both break
	i := jmpLength
	for ; i < oLen && i < jmpLength; i += jmpLength {
		if o[i] {
			break
		}
	}

	// Walk back to last jump
	for j := 0; j < jmpLength && i >= jmpLength; j++ {
		if o[i-jmpLength+j] {
			return i - jmpLength + j
		}
	}
	return -1
}
