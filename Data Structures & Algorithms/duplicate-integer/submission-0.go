func hasDuplicate(nums []int) bool {
   if len(nums) == 0 {
		return	 false;
   }

   sort.Ints(nums) // modify og array
   for i := 1; i < len(nums); i++{
	if nums[i] == nums[i-1] {
		return true
	}
   }
   return false
}
