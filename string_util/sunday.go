package string_util

func SundayIndex(s, p string) int {
	if len(p) == 0 {
		return 0
	}
	badc := make([]int, 256)
	for i := 0; i < len(badc); i++ {
		badc[i] = -1
	}

	for i, c := range p {
		badc[c] = i
	}

	si := 0
	pi := 0

	for si + len(p) <= len(s) {
		for pi < len(p) && s[si + pi] == p[pi] {
			pi++
		}

		if pi == len(p) {
			return si
		}

		if si + len(p) == len(s) {
			return -1
		}

		si += len(p) - badc[s[si+len(p)]]
		pi = 0
	}

	return -1
}
