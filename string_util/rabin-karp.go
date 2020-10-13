package string_util

// primeRK is the prime base used in Rabin-Karp algorithm.
const primeRK = 16777619

// 线性同余哈希
// hash = hash * prime + s[i]
// hashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func hashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	var pow uint32
	pow = 1
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
		pow *= primeRK
	}
	//var pow, sq uint32 = 1, primeRK
	//for i := len(sep); i > 0; i >>= 1 {
	//	if i&1 != 0 {
	//		pow *= sq
	//	}
	//	sq *= sq
	//}
	return hash, pow
}

func RabinKarpIndex(s, substr string) int {
	// Rabin-Karp search
	hashss, pow := hashStr(substr)
	n := len(substr)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	if h == hashss && s[:n] == substr {
		return 0
	}
	for i := n; i < len(s); {
		h = h*primeRK + uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashss && s[i-n:i] == substr {
			return i - n
		}
	}
	return -1
}