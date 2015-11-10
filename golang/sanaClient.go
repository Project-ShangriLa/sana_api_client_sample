//
// Copyright (c) 2015 Rirei Kuroshima
//
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const usage = "Usage: \n" +
	"    sanaClient account1 account2 account3 ...\n\n" +
	"    取得したいTwitterアカウントを指定します。"

func main() {
	if len(os.Args) == 1 {
		fmt.Println(usage)
		return
	}

	json := getSana(os.Args[1:])
	fmt.Println(json)
}

func getSana(accounts []string) string {
	url := "http://api.moemoe.tokyo/anime/v1/twitter/follower/" +
		"status?accounts="

	for i, account := range accounts {
		if i != 0 {
			url += ","
		}
		url += account
	}

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
