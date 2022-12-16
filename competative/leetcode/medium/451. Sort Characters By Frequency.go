package main

import (
	"container/heap"
	"fmt"
)

type KV struct {
	key   rune
	value int
	rank  int
}

type PQ []*KV

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].rank = i
	pq[j].rank = j
}
func (pq PQ) Less(i, j int) bool { return pq[i].value > pq[j].value }

func (pq *PQ) Push(x any) {
	n := len(*pq)
	kv := x.(*KV)
	kv.rank = n
	*pq = append(*pq, kv)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	kv := old[n-1]
	old[n-1] = nil
	kv.rank = -1
	*pq = old[0 : n-1]
	return kv
}

func frequencySort(s string) string {
	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}

	fmt.Println(m)

	queue := make(PQ, len(m))
	i := 0
	for k, v := range m {
		queue[i] = &KV{k, v, i}
		i++
	}

	heap.Init(&queue)

	res := ""

	fmt.Println(queue)

	return res
}
func main() {

}
