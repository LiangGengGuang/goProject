package main

import (
	"flag"
	"fmt"
	"os"
)

var n string

func init() {

	//等同于Usage()

	//TODO 当用户命令行输入错误时，会以状态码[status 2]提示
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)

	//TODO 当用户命令行输入错误时,会以报错形式展示 => 运行时恐慌（panic
	//flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)

	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}

	flag.CommandLine.StringVar(&n, "name", "everyone", "The greeting object.")
}

func main() {

	//自定声明
	//Usage()

	//解析
	flag.CommandLine.Parse(os.Args[1:])

	fmt.Printf("hello %s\n", n)
}

func Usage() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
}
