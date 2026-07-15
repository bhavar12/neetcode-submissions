func topKFrequent(nums []int, k int) []int {
	addFreq := make(map[int]int)
	for _, val := range nums {
		addFreq[val]++
	}

	type kv struct {
		Key   int
		Value int
	}
	var ss []kv
	for k, v := range addFreq {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	freq := []int{}
	for i := 0; i < k; i++ {
		freq = append(freq, ss[i].Key)
	}
	return freq
}