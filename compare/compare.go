package compare

type Comparer interface {
	Compare(item Comparer) int
}

type IntCompare int

func (c IntCompare) Compare(item Comparer) int {
	i := item.(IntCompare)
	if c < i {
		return -1
	}

	if c > i {
		return 1
	}

	return 0
}
