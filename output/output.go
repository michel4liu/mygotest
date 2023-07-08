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
