package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/pprof/profile"
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

	fmt.Println("id", "\t", "address", "\t", "ns", "\t", "name")
	for _, sample := range p.Sample {
		for _, location := range sample.Location {
			fmt.Println(location.ID, sample.Value[1], location.Address, location.Line[0].Function.Name, ":", location.Line[0].Line)
		}
		fmt.Println("-----------------")
	}

	sampleRaw, err := json.Marshal(p.Sample)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(pprofFilePath+".json", sampleRaw, 0644)
	if err != nil {
		panic(err)
	}
}
