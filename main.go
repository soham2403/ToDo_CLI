package main

import "fmt"

func main() {
	tds := todos{}
	tds.add("learn go")
	tds.add("uiasgfi")
	fmt.Println("main to do app")
	fmt.Println(tds)
}
