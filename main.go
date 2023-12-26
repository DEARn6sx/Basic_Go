package main //คำสั่งและpackage จะถูกทำงานในตอนเริ่มต้นในตอนที่ run program
import (
	"fmt"
	basicGo "gobasic/basicGo"
)

func main() {
	fmt.Println(basicGo.Variable())

	fmt.Println(basicGo.Sum_variadic_fnc(10, 20, 30, 50, 40, 80))

	result, check := basicGo.Return_multi_sum(1001, 600)
	fmt.Println("ผลรวม = ", result)
	fmt.Println("ผลรวมหาร = ", check)
}
