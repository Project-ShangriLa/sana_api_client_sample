//
// Copyright (c) 2015 Rirei Kuroshima
//
package sana

import (
	"errors"
	"io/ioutil"
	"net/http"
)

const usageString = "\nUsage: \n" +
	"\tfunc GetTwitter(accounts []string) (string, error)\n" +
	"\t引数で取得したツイッターアカウントが入ったスライスをを指定します。"

func GetTwitter(accounts []string) (string, error) {
	if len(accounts) == 0 {
		msg := "空のスライスが渡されています。\n"
		err := errors.New(msg + usageString)
		return "", err
	}

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
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
