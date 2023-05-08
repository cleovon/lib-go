package libgo

func Soma(values ...int) (total int) {
	for _, v := range values {
		total += v
	}

	return
}
