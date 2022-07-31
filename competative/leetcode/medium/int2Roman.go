package main

import "fmt"

func main() {
	fmt.Println(intToRoman(5))
	// fmt.Println(20 / 2)
}

type I2R struct {
	val int
	chr string
}

func intToRoman(num int) string {
	res, pending := "", false
	romanChars := []I2R{{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"}, {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"}}
	for i := 0; i < len(romanChars); {
		temp := ""
		temp, pending = checkR(&num, romanChars[i])
		res += temp
		if !pending {
			break
		}
		if num >= romanChars[i].val {
			continue
		} else {
			i++
		}
	}
	return res
}

func checkR(num *int, i2r I2R) (string, bool) {
	if *num < i2r.val {
		return "", true
	}
	// if *num < 4 {
	// 	t := ""
	// 	for i := 1; i <= *num; i++ {
	// 		t += "I"
	// 	}
	// 	return t, false
	// }
	m := *num / i2r.val
	//fmt.Println("m ", m, *num, i2r.val)
	if m > 0 {
		*num -= i2r.val
		fmt.Println("after ", *num)
		return i2r.chr, true
	} else {
		return "", *num > 0
	}
}
