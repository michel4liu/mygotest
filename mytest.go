package main

import (
	"fmt"
	"os"

	"example.com/m/v2/output"
)

var x, y int

func main() {
	var aa = 12345

	fmt.Println(&aa)

	fmt.Println(x, y)

	fmt.Printf("%v  xx  %v", x, y)

	//Print:   输出到控制台(不接受任何格式化，它等价于对每一个操作数都应用 %v)
	// fmt.Print(str)
	//Println: 输出到控制台并换行
	fmt.Println("temp")
	//Printf : 只可以打印出格式化的字符串。只可以直接输出字符串类型的变量（不可以输出整形变量和整形 等）
	fmt.Printf("%d", aa)
	//Sprintf：格式化并返回一个字符串而不带任何输出。
	// s := fmt.Sprintf("a %s", "string")
	// fmt.Printf(s)
	//Fprintf：来格式化并输出到 io.Writers 而不是 os.Stdout。
	fmt.Fprintf(os.Stderr, "an22 %s\n", "error222222")

	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
	// LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			// goto LOOP
			continue
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}

	output.OutputText(output.GetTest())

	oa, ob := output.GetTest2()
	output.OutputText(oa)
	output.OutputText(ob)

	output.GetTest3()
}
