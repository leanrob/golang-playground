package addition

func Add(nums ...int) int {
	var total int

	if len(nums) == 0 {
		println("No arguments provided")
		return 0
	}

	for _, i := range nums {
		total += i
	}
	return total
}
