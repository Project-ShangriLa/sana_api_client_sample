//
// Copyright (c) 2015 Rirei Kuroshima <rireikuroshima@icloud.com>
//
package sana

import (
	"errors"
	"io/ioutil"
	"net/http"
)

const baseurl = "http://api.moemoe.tokyo/anime/v1/twitter/follower/"

// GetLatestFollower()関数は、引数で受け取った（アニメに関する）Twitter
// アカウントをSana（紗凪）<https://github.com/Project-ShangriLa/sana_server>
// に渡し、最新のフォロワー数を取得します。
func GetLatestFollower(accounts []string) (string, error) {
	const usageString = `
Usage:
    func GetLatestFollower(accounts []string) (string, error)
    引数で取得したツイッターアカウントが入ったスライスをを指定します。`

	if len(accounts) == 0 {
		const msg = "空のスライスが渡されています。\n"
		return "", errors.New(msg + usageString)
	}

	url :=  baseurl + "status?accounts="

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

// GetFollowerHistory()関数は、引数で受け取った（アニメに関する）Twitter
// アカウントのフォロワー数の履歴をSana（紗凪）
// <https://github.com/Project-ShangriLa/sana_server>から取得します。
// 第1引数が対象のTwitterアカウント、第2引数に日時を指定します。
// この引数はunixtimestampで指定します。`
func GetFollowerHistory(args ...string) (string, error) {
	const usageString = `
Usage:
    func GetFollowerHistory(args ...string) (string, error)
    引数で取得したツイッターアカウントのフォロワー数の履歴を取得します。
    第1引数が対象のTwitterアカウント、第2引数に日時を指定します。
    この引数はunixtimestampで指定します。`

	url :=  baseurl + "history"

	if len(args) == 1 {
		url += "?account=" + args[0]
	} else if len(args) == 2 {
		url += "?account=" + args[0] +
			"&end_date=" + args[1]
	} else {
		const msg = "空のスライスが渡されています。\n"
		return "", errors.New(msg + usageString)
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
