package grade

func GetGrade(avg int) string {
	if avg >= 90 {
		return "A"
	} else if avg >= 75 {
		return "B"
	} else if avg >= 50 {
		return "C"
	}
	return "Fail"
}