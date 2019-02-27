package main

import (
	"fmt"
	"os"

	"github.com/google/pprof/profile"
	"github.com/molizz/funcpprof"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		panic("请输入待解压文件 eg. decompress xxx.pprof")
	}
	pprofFilePath := args[1]

	file, err := os.Open(pprofFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p, err := profile.Parse(file)
	if err != nil {
		panic(err)
	}

	funcs, err := funcpprof.Parse(p, true)
	if err != nil {
		panic(err)
	}

	for _, fn := range funcs {
		fmt.Println(fn.Duration, "|", fn.Name)
	}
}
