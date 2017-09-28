package leap

const testVersion = 3

func IsLeapYear(year int) bool {

	result := false

	if year % 4 == 0 && year % 100 == 0 && year % 400 == 0 {
		result = true
	}
	if year % 4 == 0 && year % 100 == 0 && year % 400 != 0 {
		result = false
	}
	if year % 4 == 0 && year % 100 != 0 && year % 400 != 0 {
		result = true
	}
	return result
}
