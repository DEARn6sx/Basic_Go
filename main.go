package main //คำสั่งและpackage จะถูกทำงานในตอนเริ่มต้นในตอนที่ run program

import "fmt" //package จัดการ input output (เกี่ยวกับการแสดงผล)

func main() {
	var number int
	fmt.Print("ป้อนตัวเลข = ")
	fmt.Scanf("%d", &number)
	fmt.Println("number is  = ", number)

	switch number {
	case 1:
		fmt.Println("เปิด บัญชี")
	case 2:
		fmt.Println("ถอนเงิน")
	default:
		fmt.Println("Error")
	}

}
