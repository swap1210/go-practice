package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

// INPUT TEMPLATE START

type CJInput struct {
	rdr         io.Reader
	lineChan    chan string
	initialized bool
}

func (mi *CJInput) start(done chan struct{}) {
	r := bufio.NewReader(mi.rdr)
	defer func() { close(mi.lineChan) }()
	for {
		line, err := r.ReadString('\n')
		if !mi.initialized {
			mi.initialized = true
			done <- struct{}{}
		}
		mi.lineChan <- strings.TrimSpace(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}
}

func (mi *CJInput) readLine() string {
	// if this is the first call, initialize
	if !mi.initialized {
		mi.lineChan = make(chan string)
		done := make(chan struct{})
		go mi.start(done)
		<-done
	}

	res, ok := <-mi.lineChan
	if !ok {
		panic("trying to read from a closed channel")
	}
	return res
}

func (mi *CJInput) readInt() int {
	line := mi.readLine()
	i, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return i
}

func (mi *CJInput) readInt64() int64 {
	line := mi.readLine()
	i, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (mi *CJInput) readInts() []int {
	line := mi.readLine()
	parts := strings.Split(line, " ")
	res := []int{}
	for _, s := range parts {
		tmp, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res = append(res, tmp)
	}
	return res
}

func (mi *CJInput) readInt64s() []int64 {
	line := mi.readLine()
	parts := strings.Split(line, " ")
	res := []int64{}
	for _, s := range parts {
		tmp, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		res = append(res, tmp)
	}
	return res
}

func (mi *CJInput) readWords() []string {
	line := mi.readLine()
	return strings.Split(line, " ")
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// INPUT TEMPLATE END

////////////////////////////////////////////////////////////////////////////////

func main() {
	f, _ := os.Open("broken_calc_input_file.txt")
	mi := CJInput{rdr: f}
	// mi := CJInput{rdr: os.Stdin}

	t := mi.readInt()
	for caseNo := 1; caseNo <= t; caseNo++ {
		fmt.Printf("Case #%d: %v\n", caseNo, brokenCalc(mi))
	}

}

func brokenCalc(mi CJInput) map[int]bool {
	sizes := mi.readInts()
	total_nums := sizes[0]
	total_ops := sizes[1]
	touches := sizes[2]
	nums := mi.readInts()
	ops := mi.readInts()
	target := mi.readInt()

	//remove equal op
	touches--
	ans := make(map[int]bool)

	fmt.Println(total_nums, total_ops, touches, nums, ops, target)
	allCombi := allCombiFunc()
	return allCombi
}

func allCombiFunc(arr []int, n int, r int, index int, data []int, i int) {

}
