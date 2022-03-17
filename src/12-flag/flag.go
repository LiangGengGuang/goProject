package main

import (
	"flag"
	"fmt"
)

//var name string
var name = flag.String("name", "everyone", "The greeting object.")

func init() {
	//flag.StringVar(&name, "name", "everyone", "The greeting object.")
}
func main() {

	flag.Parse()

	//fmt.Printf("hello %s\n", name)
	fmt.Printf("hello %s\n", *name)
}
