package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T) //T = 1 //

	for t := 1; t <= T; t++ {
		c := make([][]int, 3)
		m := [4]int{0, 0, 0, 0}

		for i := 0; i < 3; i++ {
			c[i] = make([]int, 4)
			fmt.Scan(&c[i][0], &c[i][1], &c[i][2], &c[i][3])
			if i == 0 || m[0] > c[i][0] {
				m[0] = c[i][0]
			}
			if i == 0 || m[1] > c[i][1] {
				m[1] = c[i][1]
			}
			if i == 0 || m[2] > c[i][2] {
				m[2] = c[i][2]
			}
			if i == 0 || m[3] > c[i][3] {
				m[3] = c[i][3]
			}
		}

		// fmt.Println(m[0], " ", m[1], " ", m[2], " ", m[3])
		rem := 1000000

		//max color
		for i := 0; i < 4; i++ {
			// fmt.Print(rem, m[i], "\n")
			if rem < m[i] {
				m[i] = rem
			}
			rem -= m[i]
		}

		fmt.Printf("Case #%d: ", t)
		if rem > 0 {
			fmt.Println("IMPOSSIBLE")
		} else {
			fmt.Println(m[0], " ", m[1], " ", m[2], " ", m[3])
		}
	}

}
