package addition

func Add(nums ...int) int {
	var total int
	for _, i := range nums {
		total += i
	}
	return total
}
