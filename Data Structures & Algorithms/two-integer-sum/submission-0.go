func twoSum(nums []int, target int) []int {
    idxmap:=make(map[int]int)
    for idx,val:=range nums{
      if v,ok:=idxmap[target-val];ok{
        return []int{v,idx}
      }
      idxmap[val]=idx
    }
    return []int{}
}
