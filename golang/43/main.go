package main

func main() {

}

// 中等
// 43. 字符串相乘
// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

// 运用竖式运算思想
func multiply(num1 string, num2 string) string {
	if num1 == "0" && num2 == "0" {
		return "0"
	}
	resArr := make([]int, len(num1)+len(num2))
	for i := len(num2) - 1; i >= 0; i-- {
		n2 := int(num2[i] - '0')
		for j := len(num1) - 1; j >= 0; j-- {
			n1 := int(num1[j] - '0')
			sum := n1*n1 + resArr[i+j+1]
			resArr[i+j+1] = sum % 10
			resArr[i+j] += sum / 10
		}
	}
	res := ""
	for k, v := range resArr {
		if k == 0 && v == 0 {
			continue
		}
		res += string(v + '0')
	}
	return res
}
