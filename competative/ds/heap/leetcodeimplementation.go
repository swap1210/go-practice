package main

import (
	"container/heap"
	"fmt"
)

type KV struct {
	key   rune
	value int
}

func (kv KV) String() string {
	return fmt.Sprintf("{%c, %d}", kv.key, kv.value)
}

type PQ []*KV

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq PQ) Less(i, j int) bool { return pq[i].value > pq[j].value }

func (pq *PQ) Push(x interface{}) {
	kv := x.(*KV)
	*pq = append(*pq, kv)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	if n == 0 {
		return &KV{rune('a'), 0}
	}
	kv := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return *kv
}

func frequencySort(s string) string {
	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}

	queue := &PQ{}
	heap.Init(queue)
	for k, v := range m {
		heap.Push(queue, &KV{key: k, value: v})
	}

	res := ""
	for obj := heap.Pop(queue).(KV); ; obj = heap.Pop(queue).(KV) {
		fmt.Println("Popped ", obj)
		for i := 0; i < obj.value; i++ {
			res += string(obj.key)
		}

		if len(*queue) == 0 {
			break
		}
	}

	return res
}
