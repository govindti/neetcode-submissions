func getConcatenation(nums []int) []int {
	// n := len(nums)

    // ans := make([]int, n)

	// copy(ans, nums)

	// for i:=0; i<n; i++ {
	// 	ans = append(ans, nums[i])
	// }
	// return ans


	// Even simpler:
	ans := append([]int{}, nums...)
    ans = append(ans, nums...)
    return ans

}
