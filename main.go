package main //คำสั่งและpackage จะถูกทำงานในตอนเริ่มต้นในตอนที่ run program
import "fmt" //fmt package จัดการ input output (เกี่ยวกับการแสดงผล)

type Product struct {
	name     string
	price    float64
	category string
	discount int
}

func main() {
	product1 := Product{name: "เสื้อ", price: 200, category: "ชุด", discount: 15}
	product2 := Product{name: "เกง", price: 600, category: "ชุด", discount: 50}
	fmt.Println(product1)
	fmt.Println(product1.name)
	fmt.Println(product1.price)
	fmt.Println(product1.category)
	product1.price = 500
	fmt.Println(product1.price)
	fmt.Println(product2)
}
