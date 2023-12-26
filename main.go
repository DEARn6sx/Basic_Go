package main //คำสั่งและpackage จะถูกทำงานในตอนเริ่มต้นในตอนที่ run program
import "fmt" //fmt package จัดการ input output (เกี่ยวกับการแสดงผล)

func main() {
	showmessage()
	showmessage_param("DEAR")
	total(10, 20)
	fmt.Println("ค่าจัดส่ง = ", delivery())
	fmt.Println("ผลรวมตะกร้า = ", cart(50, 80))
}

func showmessage() {
	fmt.Println("Hello Go programming")
}

func showmessage_param(name string) {
	fmt.Println("Hello ", name)
}

func total(number1 int, number2 int) {
	fmt.Println("ยอดรวม = ", number1+number2)
}

func delivery() int {
	return 50
}

func cart(number1, number2 int) int {
	return number1 + number2
}
