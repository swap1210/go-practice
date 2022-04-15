package main

import "fmt"

//UNION BY RANK
//union is fast but find connect is slow
type RankedDisjoint struct {
	root []int
	rank []int
}

func (rd *RankedDisjoint) initDisjoint(size int) {
	rd.root = make([]int, size)
	rd.rank = make([]int, size)
	for i := 0; i < size; i++ {
		rd.root[i] = i
		rd.rank[i] = 1
	}
}

func (rd *RankedDisjoint) find(toFind int) int {
	for toFind != rd.root[toFind] {
		toFind = rd.root[toFind]
	}
	return toFind
}

func (rd *RankedDisjoint) union(x int, y int) {
	rootX := rd.find(x)
	rootY := rd.find(y)
	if rootX != rootY {
		if rd.rank[rootX] > rd.rank[rootY] {
			rd.root[rootY] = rootX
		} else if rd.rank[rootX] < rd.rank[rootY] {
			rd.root[rootX] = rootY
		} else {
			rd.root[rootY] = rootX
			rd.rank[rootX] += 1
		}
	}
}

func (rd *RankedDisjoint) isConnected(x int, y int) bool {
	return rd.find(x) == rd.find(y)
}

func main() {
	fmt.Print("Disjoint set")
	given := [][2]int{{0, 1}, {1, 2}, {1, 3}, {4, 5}, {4, 6}, {1, 4}}
	l_rd := &RankedDisjoint{}
	l_rd.initDisjoint(7)
	for _, node := range given {
		l_rd.union(node[0], node[1])
	}
	fmt.Println(l_rd.root)
	fmt.Println("6 and 1 connected is ", l_rd.isConnected(6, 1))
}
