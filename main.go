package main //คำสั่งและpackage จะถูกทำงานในตอนเริ่มต้นในตอนที่ run program

import "fmt" //package จัดการ input output (เกี่ยวกับการแสดงผล)

//func main = fnc ที่ go มองอันดับแรก
func main() {
	fmt.Println("Hello Go Programming") //แสดงผลข้อความและขึ้นบรรทัดใหม่

	//นิยามตัวแปร
	var name string = "DEAR"
	lastname := "ZAZA"
	var age int = 26
	var score float32 = 95.8
	var isPass bool = true
	fmt.Println("My name is", name)
	fmt.Println("My lastname is", lastname)
	fmt.Println("My age is", age)
	fmt.Println("My score is", score)
	fmt.Println("My testing is", isPass)

	//data type and value
	address := "roi et"
	fmt.Printf("My value address is %v\n", address)
	fmt.Printf("My data type address is %T\n", address)

	//math operator
	// var num1 int = 10
	// var num2 int = 3
	var num1, num2 = 10, 3
	fmt.Println("ผลบวก = ", num1+num2)
	fmt.Println("ผลลบ = ", num1-num2)
	fmt.Println("ผลคูณ = ", num1*num2)
	fmt.Println("ผลหาร = ", num1/num2)
	fmt.Println("ผลหานเอาเศษ = ", num1%num2)
}
