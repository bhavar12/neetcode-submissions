func findMaxConsecutiveOnes(nums []int) int {
    maxCount:=0
    currCount:=0
    for _, val:=range nums{
       if val==1{
        currCount++
       }else{
 if currCount > maxCount{
        maxCount = currCount
      }
       currCount = 0
       }
    
    }
    if currCount > maxCount{
      return currCount
    }
    return maxCount
}
