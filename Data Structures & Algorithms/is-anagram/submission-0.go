func isAnagram(s string, t string) bool {
  if len(s) != len(t){
    return false
  }
  freqS:=make(map[rune]int)
  freqT:=make(map[rune]int)
  for _,v:=range s{
    if _,ok:=freqS[v];ok{
      freqS[v]+=1
    }else{
      freqS[v]=1
    }
  }

  for _,v:=range t{
    if _,ok:=freqT[v];ok{
      freqT[v]=freqT[v]+1
    }else{
      freqT[v]=1
    }
  }
  for key,val:=range freqS{
    if count,ok:=freqT[key];ok{
      if count != val{
        return false
      }
      continue
    }
    return false
  }
  return true
}
