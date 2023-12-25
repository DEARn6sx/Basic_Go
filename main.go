package main //คำสั่งและpackage จะถูกทำงานในตอนเริ่มต้นในตอนที่ run program
import "fmt"

//package จัดการ input output (เกี่ยวกับการแสดงผล)

func main() {
	var numbers [3]int
	fmt.Println(numbers)

	numbers[0] = 100
	numbers[1] = 200
	numbers[2] = 300
	fmt.Println(numbers)

	var numbers2 [3]int = [3]int{400, 500, 600}
	fmt.Println(numbers2)

	numbers3 := [3]int{700, 800, 'a'}
	fmt.Println(numbers3)

	numbers4 := [...]int{700, 800, 'a', 1, 2, 3, 4, 5, 6}
	fmt.Println(numbers4)

	size := len(numbers4)
	fmt.Println(size)

}
