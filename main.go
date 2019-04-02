package main

import (
	"concurrency_run/tools"
	"fmt"
	"math"
	"runtime"
)

// 平行處理參考這邊
// https://larrylu.blog/golang-goroutine-parallel-processing-d382e6d34a14

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := tools.ReadLines("./csv/file.txt")
	check(err)

	nCPU := runtime.NumCPU()
	maxNum := float64(len(dat)) / float64(nCPU)
	nNum := int(math.Round(maxNum))
	// idx := 0
	result := ""
	x := ""
	ch := make(chan string)

	for loop := 0; loop < nNum; loop++ {
		for i := 0; i < nCPU; i++ {
			go func() {
				if len(dat) > 0 {
					x, dat = dat[len(dat)-1], dat[:len(dat)-1]
					fmt.Println(tools.ReadLines(x))
					ch <- x
				} else {
					ch <- ""
				}
			}()
		}

		for i := 0; i < nCPU; i++ {
			result = <-ch
			fmt.Println(result)
		}
	}
}
