package main

import (
	"fmt"

	v "github.com/libotony/pkgexample/v2/version"
	"github.com/vechain/thor/thor"
)

func main() {
	addr, err := thor.ParseAddress("0xdb54b4af8e040ce4f88b15afd6df5c00db42cfcf")
	if err == nil {
		fmt.Println(addr)
	}
	fmt.Printf("pkg example version %s\n", v.Version())
}
