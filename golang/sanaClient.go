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
	"strings"
)

const usage = "sanaClient:\n" +
	"    -a string\n" +
	"        取得したいTwitterアカウント"

func main() {
	var accounts string
	flag.StringVar(&accounts, "a", "", "取得したいTwitterアカウント")
	flag.Parse()

	if accounts == "" {
		fmt.Println(usage)
		return
	}

	json := getSana(strings.Split(accounts, ","))
	fmt.Println(json)
}

func getSana(acount []string) string {
	url := "http://api.moemoe.tokyo/anime/v1/twitter/follwer/" +
		"status?accounts="

	for i, a := range acount {
		if i != 0 {
			url += ","
		}
		url += a
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
