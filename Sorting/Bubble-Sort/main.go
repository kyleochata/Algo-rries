package main

/*
every iteration of bubble sort will place the largest value at the end
O(n^2)
[0, ...N] compare i to i+1. don't go past N-1
for i...n (do n amt of times)
for j ..n-1-i (after each iteration, the largest item was placed to the R. can stop the iteration i+1 times early)
if (arr[i] > arr[j]){swap}
*/
func bubble_sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
