package main //คำสั่งและpackage จะถูกทำงานในตอนเริ่มต้นในตอนที่ run program
import "fmt"

//package จัดการ input output (เกี่ยวกับการแสดงผล)

func main() {

	//map
	country_map := map[string]string{"th": "thailand", "jpn": "japan"}
	fmt.Println(country_map)
	fmt.Println(country_map["th"])
	fmt.Println(country_map["jpn"])

	value, check := country_map["th"]
	//check = ตรวจสอบว่ามี "th" ไหม
	if check {
		fmt.Println((value))
	} else {
		fmt.Println("not found")
	}
}
