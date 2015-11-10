//
// Copyright (c) 2015 Rirei Kuroshima <rireikuroshima@icloud.com>
//
package sana

import (
	"fmt"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
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
		utime, err := changeTime(args[1])
		if err != nil {
			const msg = "日時のフォーマットが違います。\n"
			return "", errors.New(msg + usageString)
		}
		url += "?account=" + args[0] +
			"&end_date=" + utime
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

// YYYY-MM-DDの形式で受け取った文字列をUnixTimestampに変換する。
func changeTime(t string) (string, error) {
	var year, month, day int
	var err error

	str := strings.Split(t, "-")
	if len(str) != 3 {
		return "", errors.New("YYYY-MM-DDで指定してください。")
	}
	year, err = strconv.Atoi(str[0])
	month, err = strconv.Atoi(str[1])
	day, err = strconv.Atoi(str[2])
	if err != nil {
		return "", err
	}
	if year > time.Now().Format("2006") {
		return "", errors.New("年月日を正しく指定してください。")
	}
	if month < 1 || 12 < month {
		return "", errors.New("年月日を正しく指定してください。")
	}
	if day < 1 || 31 < day {
		return "", errors.New("年月日を正しく指定してください。")
	}

	convertTime := time.Date(year, time.Month(month), day,
		23, 59, 59, 0, time.Local)
	convertTime = convertTime.Add(time.Duration(1) * time.Second)

	return string(convertTime.Unix()), nil
}
