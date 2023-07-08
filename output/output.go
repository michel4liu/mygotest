package output

import "fmt"

func OutputText(s string) {
	outputText2(s)
}

func outputText2(s string) {
	fmt.Println(s)
}

func GetTest() string {
	return "你好呀~"
}

func GetTest2() (string, string) {
	return "第一个结果", "第二个结果"
}

func GetTest3() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}
