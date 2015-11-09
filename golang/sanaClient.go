//
// Copyright (c) 2015 Rirei Kuroshima
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const usage = "sanaClient:\n" +
	"    -a string\n" +
	"        取得したいTwitterアカウント"

func main() {
	const baseurl = "http://api.moemoe.tokyo/anime/v1/twitter/follwer/"

	var accounts string
	flag.StringVar(&accounts, "a", "", "取得したいTwitterアカウント")
	flag.Parse()

	if accounts == "" {
		fmt.Println(usage)
		return
	}

	url := baseurl + "status?accounts=" + accounts

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	result := string(body)

	fmt.Println(result)
}
