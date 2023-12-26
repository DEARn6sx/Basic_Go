package basicGo

func Return_multi_sum(num1, num2 int) (int, string) {
	total := num1 + num2
	status := ""
	if total%2 == 0 {
		status = "เลขคู่"
	} else {
		status = "เลขคี่"
	}
	return total, status
}