package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(sl []string) FreqMap {
	op := FreqMap{}
	c := make(chan FreqMap)
	for _, s := range sl {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}
	for range sl {
		for k, v := range <-c {
			op[k] += v
		}
	}
	return op
}
