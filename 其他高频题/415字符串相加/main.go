package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	if n1 > n2 {
		n1, n2 = n2, n1
		num1, num2 = num2, num1
	}
	var patch uint8
	res := make([]byte, 0, n2)
	// num1是短的，num1 + num2 + 进位
	for i := n2 - 1; i >= 0; i-- {
		sum := (num2[i] - '0') + patch
		if j := i - (n2 - n1); j >= 0 {
			sum += num1[j] - '0'
		}
		//var unit uint8
		//if sum >= 10 {
		//	patch = 1
		//	unit = sum - 10
		//} else {
		//	patch = 0
		//	unit = sum
		//}
		// 这里求个数和十位可以用除法和取余代替
		unit := sum % 10
		patch = sum / 10
		res = append(res, unit+'0')
	}
	// num2 + 进位
	if patch == 1 {
		res = append(res, '1')
	}
	l, r := 0, len(res)-1
	for l <= r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return string(res)
}

func main() {
	fmt.Println(addStrings("123", "947"))
}
