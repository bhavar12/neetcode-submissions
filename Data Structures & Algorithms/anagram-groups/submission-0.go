func groupAnagrams(strs []string) [][]string {
res:=make(map[string][]string)
for _,s:= range strs{
  sortedS:=sortedString(s)
  res[sortedS] = append(res[sortedS],s)
}
var result [][]string
 for _,group:=range res{
  result = append(result,group)
 }
 return result
}

func sortedString(s string) string{
  chars:=[]rune(s)
  sort.Slice(chars,func(i,j int) bool{
    return chars[i] < chars[j]
  })
  return string(chars)
}
