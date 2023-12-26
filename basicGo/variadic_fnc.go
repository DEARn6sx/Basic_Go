package basicGo

func Sum_variadic_fnc(num ...int) int {
	total := 0
	for _, value := range num {
		total += value
	}
	return total
}
