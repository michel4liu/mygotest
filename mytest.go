package main

import (
	"encoding/json"
	"fmt"
	"os"

	"example.com/m/v2/output"
)

var x, y int

func main() {

	const ca, cb int = 1, 2

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

	// 创建一个新的结构体
	//fmt.Println(StrutcTest2{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// // 也可以使用 key => value 格式
	// fmt.Println(StrutcTest{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	// // 忽略的字段为 0 或 空
	// fmt.Println(StrutcTest{title: "Go 语言", author: "www.runoob.com"})

	//一维数组
	var arr_1 [5]int
	fmt.Println(arr_1)

	var arr_2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_2)

	arr_3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_3)

	arr_4 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr_4)

	arr_5 := [5]int{0: 3, 1: 5, 4: 6}
	fmt.Println(arr_5)

	//二维数组
	var arr_6 = [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_6)

	arr_7 := [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_7)

	arr_8 := [...][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {0: 3, 1: 5, 4: 6}}
	fmt.Println(arr_8)

	var arr = [...]int{1, 2, 3, 4, 5}
	modifyArr(&arr)
	fmt.Println(arr)

	var sli_1 []int //nil 切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1), cap(sli_1), sli_1)

	var sli_2 = []int{} //空切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1), cap(sli_2), sli_2)

	var sli_3 = []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_3), cap(sli_3), sli_3)

	sli_4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_4), cap(sli_4), sli_4)

	var sli_5 []int = make([]int, 5, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_5), cap(sli_5), sli_5)

	sli_6 := make([]int, 5, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_6), cap(sli_6), sli_6)

	fmt.Println("------------------------------------------------")

	// sli := [] int {1, 2, 3, 4, 5, 6}
	// fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	// fmt.Println("sli[1] ==", sli[1])
	// fmt.Println("sli[:] ==", sli[:])
	// fmt.Println("sli[1:] ==", sli[1:])
	// fmt.Println("sli[:4] ==", sli[:4])

	// fmt.Println("sli[0:3] ==", sli[0:3])
	// slix:=sli[0:3]
	// fmt.Println(cap(slix))
	// fmt.Printf("len=%d cap=%d slice=%v\n",len(sli[0:3]),cap(sli[0:3]),sli[0:3])

	// fmt.Println("sli[0:3:4] ==", sli[0:3:4])
	// fmt.Printf("len=%d cap=%d slice=%v\n",len(sli[0:3:4]),cap(sli[0:3:4]),sli[0:3:4])

	fmt.Println("------------------------------------------------")
	// sli := [] int {4, 5, 6}
	// fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	// sli = append(sli, 7)
	// fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	// sli = append(sli, 8)
	// fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	// sli = append(sli, 9)
	// fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	// sli = append(sli, 10)
	// fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	sli := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	//删除尾部 2 个元素
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[:len(sli)-2]), cap(sli[:len(sli)-2]), sli[:len(sli)-2])

	//删除开头 2 个元素
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[2:]), cap(sli[2:]), sli[2:])

	//删除中间 2 个元素
	sli = append(sli[:3], sli[3+2:]...)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	//匿名结构体
	p4 := struct {
		Name string `json:"名称"`
		Age  int    `json:"-"`
		sex  bool
	}{Name: "匿名222222", Age: 33}
	fmt.Println("p4 =", p4)

	//序列化
	jsons, errs := json.Marshal(p4)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data :", string(jsons))

	var p1 = make(map[int]string)
	p1[1] = "Tom"
	p1[1] = "Tom333"
	p1[2] = "Tom2"
	fmt.Println("p1 :", p1)

	var p2 map[int]string = map[int]string{}
	p2[1] = "Tom"
	fmt.Println("p2 :", p2)

	var p3 map[int]string = make(map[int]string)
	p3[1] = "Tom"
	fmt.Println("p3 :", p3)

	// p4 := map[int]string{}
	// p4[1] = "Tom"
	// fmt.Println("p4 :", p4)

	p5 := make(map[int]string)
	p5[1] = "Tom"
	fmt.Println("p5 :", p5)

	p6 := map[int]string{
		1: "Tom",
	}
	fmt.Println("p6 :", p6)

	fmt.Println("------------------------------------------------")
	person := [3]string{"Tom", "Aaron", "John"}
	fmt.Printf("len=%d cap=%d array=%v\n", len(person), cap(person), person)

	fmt.Println("")

	//循环
	for k, v := range person {
		fmt.Printf("person[%d]: %s\n", k, v)
	}

	fmt.Println("")

	for i := range person {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}

	fmt.Println("")

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}

	fmt.Println("")

	//使用空白符
	for _, name := range person {
		fmt.Println("name :", name)
	}

	fmt.Println("------------------------------------------------")

	switchi := 3
	fmt.Printf("当 i = %d 时：\n", switchi)

	switch switchi {
	case 1:
		fmt.Println("输出 i =", 1)
	case 2:
		fmt.Println("输出 i =", 2)
	case 3:
		fmt.Println("输出 i =", 3)
		fallthrough
	case 4, 5, 6:
		fmt.Println("输出 i =", "4 or 5 or 6")
	default:
		fmt.Println("输出 i =", "xxx")
	}

	fmt.Println("------------------------------------------------")
	var numbers = make([]int, 0, 5)
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4, 5, 6, 7, 9, 11, 8)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func modifyArr(a *[5]int) {
	a[1] = 20
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
