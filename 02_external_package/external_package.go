package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}

//  Run:
// 		go mod init external_package
//    go run external_package.go
