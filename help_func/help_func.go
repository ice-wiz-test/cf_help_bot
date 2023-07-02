package help_func

func Max(data []int) int {
	max := -10000000
	for i := 0; i < len(data); i++ {
		if max < data[i] {
			max = data[i]
		}
	}
	return max
}
