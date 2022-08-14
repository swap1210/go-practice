package main

import "fmt"

func main() {
	x := []string{"abc", "defg", "xyzs"}
	fmt.Println(allCombo(x))
}

func allCombo(in []string) []string {
	q := []string{""}

	res := []string{}

	for len(q) != 0 {
		fmt.Println("q is ", q)
		s := q[len(q)-1]
		q = q[:len(q)-1]

		//this length can be less but can't be more
		if len(s) == len(in) {
			res = append(res, s)
		} else {
			for _, v := range in[len(s)] {
				q = append(q, s+string(v))
			}
		}
	}

	return res
}
