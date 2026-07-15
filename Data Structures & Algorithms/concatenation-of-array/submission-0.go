func getConcatenation(nums []int) []int {
	if len(nums) == 0{
		return []int{}
	}
    ans:=make([]int,len(nums)*2)
	lenght:=len(nums)
	for i,n:=range nums{
     ans[i] = n
	 ans[i+lenght]=n
	}
	return ans
}
