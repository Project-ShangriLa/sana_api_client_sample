package sana

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetTwitter(accounts []string) string {
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
