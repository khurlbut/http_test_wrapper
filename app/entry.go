package main

import (
	"fmt"
	en "kdh.com/greet"
	de "kdh.com/greet/de"
)

var integers [10]int

func init() {
	fmt.Println("app/entry.go ==> init()")

	for i := 0; i < 10; i++ {
		integers[i] = i
	}
}

var myVersion = fetchVersion()

func main() {
	fmt.Println("app/fetch-version.go ==> fetchVersion()")
	fmt.Println("version ==> ", myVersion)

	fmt.Println(integers)

	fmt.Println("kdh.com/app/entry.g ==> main()")
	fmt.Println(en.Morning)
	fmt.Println(de.Morning)
}
