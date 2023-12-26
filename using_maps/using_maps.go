package using_maps

func Freq(names []string) map[string]int {
	freq := make(map[string]int)
	for _, name := range names {
		freq[name]++
	}
	return freq
}
