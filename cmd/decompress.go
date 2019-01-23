/*
解压gzip并输出
*/
package main

import (
	"fmt"
	"os"

	"github.com/google/pprof/profile"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		panic("请输入待解压文件 eg. decompress xxx.pprof")
	}
	comFilePath := args[1]

	file, err := os.Open(comFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p, err := profile.Parse(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("id", "\t", "ns", "\t", "name")
	for _, sample := range p.Sample {
		for _, location := range sample.Location {
			fmt.Println(location.ID, sample.Value[1], location.Line[0].Function.Name)
		}
		fmt.Println("-----------------")
	}
}
