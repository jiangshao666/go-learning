package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func logicCode() {
	var ch chan int
	for {
		select {
		case v := <-ch: //阻塞
			fmt.Println("receive from ch", v)
		default:
		}
	}
}

func main() {
	var isCpuProf bool
	var isMemProf bool

	flag.BoolVar(&isCpuProf, "cpu", false, "开启cpu调试")
	flag.BoolVar(&isMemProf, "mem", false, "开启内存调试")
	flag.Parse()

	if isCpuProf {
		f1, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu profile failed, err:%v", err)
			return
		}
		pprof.StartCPUProfile(f1)
		defer func() {
			pprof.StopCPUProfile()
			f1.Close()
		}()
	}

	for i := 0; i < 6; i++ {
		go logicCode()
	}
	time.Sleep(time.Second * 20)
	if isMemProf {
		f2, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem profile failed, err:%v", err)
			return
		}
		pprof.WriteHeapProfile(f2)
		f2.Close()
	}
}
