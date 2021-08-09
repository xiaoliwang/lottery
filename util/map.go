package util

import "lottery/lottery"

func NewStringInt(src map[string]int) (map[string]int, error) {
	dist := make(map[string]int)
	for k, v := range src {
		dist[k] = v
	}
	return dist, nil
}

func Keys(m lottery.Possibility) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
