package string_util

func BMIndex(s, p string) int {
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

	suffix := make([]int, len(p))
	suffix[len(p) - 1] = 0
	for i := 0; i < len(p) - 1; i++ {
		k := 0
		for i >= k && p[i - k] == p[len(p) - 1 - k] {
			k++
		}
		suffix[i] = k
	}

	//fmt.Println(suffix)
	bm := make([]int, len(p))
	for i := 0; i < len(p); i++ {
		bm[i] = len(p)
	}
	for i := len(p) - 2; i >= 0; i-- {
		if suffix[i] == i+1 {
			for j:=0; j < len(p) - 1; j++ {
				bm[j] = len(p) - 1 - i
			}
			break
		}
	}

	for i := 0; i < len(p); i++ {
		for j := 1; j <= suffix[i]; j++ {
			{
				bm[len(p) - 1 - j] = len(p) - 1 - i
			}
		}
	}

	bm[len(p)-1]=1

	//fmt.Println(bm)

	si := 0
	for si + len(p) <= len(s) {
		pi := len(p) - 1
		for pi >= 0 && s[si + pi] == p[pi] {
			pi--
		}
		if pi < 0 {
			return si
		}

		mv := pi - badc[s[si + pi]]
		mv = -1
		if mv < bm[pi] {
			mv = bm[pi]
		}

		if mv <= 0 {
			mv = 1
		}
		si += mv
	}

	return -1
}
