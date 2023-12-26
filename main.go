package main //คำสั่งและpackage จะถูกทำงานในตอนเริ่มต้นในตอนที่ run program
import "fmt" //fmt package จัดการ input output (เกี่ยวกับการแสดงผล)

func main() {
	result := sum(1001, 600,50,60,80,11111)
	fmt.Println("ผลรวม = ", result)
}

func sum(num... int) (int) {
	total := 0
	for _, value := range num {
		total += value
	}
	return total
}
