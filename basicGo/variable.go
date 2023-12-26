package basicGo

func Variable() (string, string, string, int, float32, bool) {
	//fmt.Println("Hello Go Programming") //แสดงผลข้อความและขึ้นบรรทัดใหม่
	//นิยามตัวแปร
	var name string = "DEAR"
	lastname := "ZAZA"
	var age int = 26
	var score float32 = 95.8
	var isPass bool = true
	// fmt.Println("My name is", name)
	// fmt.Println("My lastname is", lastname)
	// fmt.Println("My age is", age)
	// fmt.Println("My score is", score)
	// fmt.Println("My testing is", isPass)
	return "Hello Go Programming", name, lastname, age, score, isPass
}
