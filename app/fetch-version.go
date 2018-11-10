package main

import (
	"fmt"
	version "kdh.com/version"
)

func init() {
	fmt.Println("app/fetch-version.go ==> init()")
}

func fetchVersion() string {
	fmt.Println("app/fetch-version.go ==> fetchVersion()")
	return version.Version
}
