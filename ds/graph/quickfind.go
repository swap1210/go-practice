package main

import "fmt"

//QUICK FIND
//union is slow but find connect is fast
	var root []int

	func initDisjoint(size int){
		root = make([]int, size)
		for i:=0;i<size;i++{
			root[i] = i
		}
	}

	func find(toFind int) int{
		return root[toFind]
	}

	func union(x int,y int){
		rootX := find(x)
		rootY := find(y)

		if(rootX != rootY){
			for i:=0;i<len(root);i++{
				if(root[i] == rootY){
					root[i] = rootX
				}
			}
		}
	}

	func isConnected(x int,y int) bool{
		return root[x] == root[y]
	}

func main(){
	fmt.Print("Root set")
	given := [...][2]int{{0,1},{1,3},{1,4},{2,5},{2,6},{0,2}}
	initDisjoint(7)
	for _,node := range(given){
		union(node[0],node[1])
	}
	fmt.Print(root)
	fmt.Print("\nFound 4 at ",find(4))
	fmt.Print("\n6 and 1 connected is ",isConnected(6,1))
}
