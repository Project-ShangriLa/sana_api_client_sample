//
// Copyright (c) 2015 Rirei Kuroshima
//
package main

import (
	"fmt"
	"os"
	"./sana"
)

const usage = "Usage: \n" +
	"    sanaClient account1 account2 account3 ...\n\n" +
	"    取得したいTwitterアカウントを指定します。"

func main() {
	if len(os.Args) == 1 {
		fmt.Println(usage)
		return
	}

	json := sana.GetTwitter(os.Args[1:])
	fmt.Println(json)
}
