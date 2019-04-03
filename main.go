package main

import (
	"concurrency_run/tools"
	"fmt"
	"math"
	"os/exec"
	"runtime"
	"sync"
)

// 平行處理參考這邊
// https://larrylu.blog/golang-goroutine-parallel-processing-d382e6d34a14

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func execCmdBy(data []string) {
	cmd := exec.Command("echo", "Hello")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	dat, err := tools.ReadLines("./csv/file.txt")
	check(err)
	
	var mutex = &sync.Mutex{}

	nCPU := runtime.NumCPU()
	maxNum := float64(len(dat)) / float64(nCPU)
	nNum := int(math.Round(maxNum))
	x := ""
	ch := make(chan string)

	for loop := 0; loop < nNum; loop++ {
		for i := 0; i < nCPU; i++ {
			go func() {
				if len(dat) > 0 {
					mutex.Lock()
					x, dat = dat[len(dat)-1], dat[:len(dat)-1]
					content, _ := tools.ReadLines(x)
					execCmdBy(content)
					mutex.Unlock()
					ch <- x
				} else {
					ch <- ""
				}
			}()
		}

		for i := 0; i < nCPU; i++ {
			<-ch
		}
	}
}
