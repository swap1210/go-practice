package main

import "fmt"

//QUICK UNION
//union is fast but find connect is slow
	var root []int

	func initDisjoint(size int){
		root = make([]int, size)
		for i:=0;i<size;i++{
			root[i] = i
		}
	}

	func find(toFind int) int{
		for;toFind != root[toFind];{
			toFind = root[toFind]
		}
		return toFind
	}

	func union(x int,y int){
		rootX := find(x)
		rootY := find(y)
		if rootX != rootY{
			root[rootY] = rootX
		}
	}

	func isConnected(x int,y int) bool{
		return find(x) == find(y)
	}

func main(){
	fmt.Print("Disjoint set")
	given := [...][2]int{{0,1},{1,2},{1,3},{4,5},{4,6},{1,5}}
	initDisjoint(7)
	for _,node := range(given){
		union(node[0],node[1])
	}
	fmt.Print(root)
	fmt.Print("\n4 and 5 connected is ",isConnected(4,5))
}
