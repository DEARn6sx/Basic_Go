package main //คำสั่งและpackage จะถูกทำงานในตอนเริ่มต้นในตอนที่ run program

import "fmt" //package จัดการ input output (เกี่ยวกับการแสดงผล)

// if else
func main() {
	var score int
	fmt.Print("ป้อนคะแนน = ")
	fmt.Scanf("%d", &score)
	fmt.Println("score is  = ", score)

	if score >= 50 {
		fmt.Println("Pass")
	} else if score < 50 {
		fmt.Println("Fail")
	} else {
		fmt.Println("Error")
	}
}
