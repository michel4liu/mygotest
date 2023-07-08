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
