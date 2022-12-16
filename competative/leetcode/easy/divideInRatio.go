package main

func main() {
	num := 5
	print(divide(num))
}

func divide(num int) (int, int, int) {
	x, y := 1, 0

	for ; ; x++ {
		if (num - 2*x) > 0 {
			y = num - 2*x
		} else {
			break
		}
	}
	return x, y, x
}
