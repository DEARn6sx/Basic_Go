package main //คำสั่งและpackage จะถูกทำงานในตอนเริ่มต้นในตอนที่ run program

import "fmt" //package จัดการ input output (เกี่ยวกับการแสดงผล)

// scanf
func main() {
	var name string
	fmt.Print("ป้อนชื่อ = ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hi = ", name)
}
