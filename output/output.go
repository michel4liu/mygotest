package output

import "fmt"

func OutputText() {
	outputText2("这是一段文字。")
}

func outputText2(s string) {
	fmt.Println(s)
}
