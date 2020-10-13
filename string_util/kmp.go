package string_util

func KMPIndex(s, p string) int {
	if len(p) == 0 {
		return 0
	}
	next := make([]int, len(p))
	next[0] = -1
	for i := 1; i < len(p); i++ {
		next[i] = -1
		j := next[i-1]
		for j != -1 && p[i] != p[j + 1] {
			j = next[j]
		}
		if p[i] == p[j+1] {
			next[i] = j + 1
		}
	}

	for i := len(p) - 1; i > 0; i-- {
		next[i] = next[i-1] + 1
	}


	for i := 0; i < len(p); i++ {
		if next[i] >= 0 && p[next[i]] == p[i] {
			next[i] = next[next[i]]
		}
	}

	si := 0
	pi := 0

	for si < len(s) {
		for si < len(s) && pi < len(p) && s[si] == p[pi] {
			si++
			pi++
		}

		if pi >= len(p) {
			return si - len(p)
		}

		pi = next[pi]
		if pi < 0 {
			si++
			pi++
		}

		//if pi == 0 {
		//	si++
		//} else {
		//	pi = next[pi-1] + 1
		//}

	}

	return -1
}
