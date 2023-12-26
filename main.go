package main //คำสั่งและpackage จะถูกทำงานในตอนเริ่มต้นในตอนที่ run program
import "fmt" //fmt package จัดการ input output (เกี่ยวกับการแสดงผล)

func main() {
	result, check := sum(1001, 600)
	fmt.Println("ผลรวม = ", result)
	fmt.Println("ผลรวมหาร = ", check)
}

func sum(num1, num2 int) (int, string) {
	total := num1 + num2
	status := ""
	if total%2 == 0 {
		status = "เลขคู่"
	} else {
		status = "เลขคี่"
	}
	return total, status
}
