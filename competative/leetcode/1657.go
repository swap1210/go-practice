package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [2]int{2, 5}
	b := [2]int{2, 5}
	if a == b {
		fmt.Println("match")
	}
	fmt.Println(closeStrings("abbzzca", "babzzcz"))
}

func closeStrings(word1 string, word2 string) bool {
	counter := func(word string) (keys, vals [26]int) {
		for i := range word {
			keys[word[i]-'a'] = 1
			vals[word[i]-'a'] += 1
		}
		sort.Ints(vals[:])
		return keys, vals
	}
	keys1, vals1 := counter(word1)
	keys2, vals2 := counter(word2)
	return keys1 == keys2 && vals1 == vals2
}

// my simple solution
/*
func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2){
        return false
    }

    m1 := make(map[byte]int)
    m2 := make(map[byte]int)
    for i:=0;i<len(word1);i++ {
        m1[word1[i]]++
        m2[word2[i]]++
    }

    a1 := make([]int,len(m1))
    a2 := make([]int,len(m2))
    i:= 0
    for k,_ := range m1{
        fmt.Println("For ",string(k))
        if _,ok := m2[k];!ok{
            return false
        }
        a1[i] = m1[k]
        a2[i] = m2[k]
        i++
    }
    sort.Ints(a1)
    sort.Ints(a2)

    return Equals(a1,a2)
}

func Equals(a,b []int) bool{
    for i,_ := range a{
        if a[i] != b[i]{
            return false
        }
    }
    return true
}
*/
