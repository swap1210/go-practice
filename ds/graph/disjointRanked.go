package main

import "fmt"

//UNION BY RANK
//union is fast but find connect is slow
var root []int
var rank []int

func initDisjoint(size int) {
	root = make([]int, size)
	rank = make([]int, size)
	for i := 0; i < size; i++ {
		root[i] = i
		rank[i] = 1
	}
}

func find(toFind int) int {
	for toFind != root[toFind] {
		toFind = root[toFind]
	}
	return toFind
}

func union(x int, y int) {
	rootX := find(x)
	rootY := find(y)
	if rootX != rootY {
		if rank[rootX] > rank[rootY] {
			root[rootY] = rootX
		} else if rank[rootX] < rank[rootY] {
			root[rootX] = rootY
		} else {
			root[rootY] = rootX
			rank[rootX] += 1
		}
	}
}

func isConnected(x int, y int) bool {
	return find(x) == find(y)
}

func main() {
	fmt.Print("Disjoint set")
	given := [][2]int{{0, 1}, {1, 2}, {1, 3}, {4, 5}, {4, 6}, {1, 4}}
	initDisjoint(7)
	for _, node := range given {
		union(node[0], node[1])
	}
	fmt.Print(root)
	fmt.Print("\n6 and 1 connected is ", isConnected(6, 1))
}
