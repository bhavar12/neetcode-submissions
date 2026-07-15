func productExceptSelf(nums []int) []int {
if len(nums) == 0{
	return []int{}
}
powerarr:=make([]int,len(nums))
for i:=0;i<len(nums);i++{
	prod:=1
	for j:=0;j<len(nums);j++{
		if i !=j{
		prod=prod*nums[j]
		}
	}
	powerarr[i]=prod
}
return powerarr
}
