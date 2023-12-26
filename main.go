package main //คำสั่งและpackage จะถูกทำงานในตอนเริ่มต้นในตอนที่ run program
import "fmt"

//package จัดการ input output (เกี่ยวกับการแสดงผล)

func main() {

	for i := 0; i < 10; i++ {

		if i == 5 {
			break
		}
		if i == 2 {
			continue
		}
		fmt.Println("Hello Go", i)
	}
	fmt.Println("end")
}
