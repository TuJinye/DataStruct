package main

import (
	"fmt"
)

func main() {
	// 定义字符串变量
	var str1 string = "Hello, World!"
	fmt.Println("str1:", str1)

	// 获取字符串长度
	length := len(str1)
	fmt.Println("Length of str1:", length)

	// 遍历字符串
	fmt.Print("Characters in str1:")
	for i := 0; i < length; i++ {
		fmt.Printf(" %c", str1[i])
	}
	fmt.Println()

	// 拼接字符串
	str2 := " How are you?"
	str3 := str1 + str2
	fmt.Println("str3:", str3)

	// 判断字符串相等
	if str1 == "Hello, World!" {
		fmt.Println("str1 is equal to 'Hello, World!'")
	}

	// 判断字符串包含子串
	if contains := Contains(str3, "Hello"); contains {
		fmt.Println("str3 contains 'Hello'")
	}

	// 截取字符串
	substr := str3[7:12]
	fmt.Println("Substring of str3:", substr)

	// 将字符串转换为字节数组
	byteArray := []byte(str3)
	fmt.Println("Byte array:", byteArray)

	// 将字节数组转换为字符串
	str4 := string(byteArray)
	fmt.Println("str4:", str4)
}

// 判断字符串是否包含子串
func Contains(s string, substr string) bool {
	return len(substr) <= len(s) && s[0:len(substr)] == substr
}
