//
// Copyright (c) 2015 Rirei Kuroshima <rireikuroshima@icloud.com>
//
package main

import (
	"fmt"
	"log"
	"os"
	"./sana"
)

func main() {
	if json, err := sana.GetLatestFollower(os.Args[1:]); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(json)
	}

	if json, err := sana.GetFollowerHistory("usagi_anime"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(json)
	}

	json, err := sana.GetFollowerHistory("usagi_anime", "1407562541")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(json)
	}
}
