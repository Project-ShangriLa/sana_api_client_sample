//
// Copyright (c) 2015 Rirei Kuroshima <rireikuroshima@icloud.com>
//
package sana

import (
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

	return getJson(url)
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

	return getJson(url)
}

// YYYY-MM-DDの形式で受け取った文字列をUnixTimestampに変換する。
func changeTime(t string) (string, error) {
	var year, month, day, hour, minute, second int
	var err error

	// YYYY-MM-DDとHH:ii:ssに分割する
	date := strings.Split(t, " ")
	if len(date) != 2 {
		return "", errors.New("YYYY-MM-DD HH:ii:ssで指定してください。")
	}

	// 年月日の処理
	str1 := strings.Split(date[0], "-")
	if len(str1) != 3 {
		return "", errors.New("年月日はYYYY-MM-DDで指定してください。")
	}
	year, err = strconv.Atoi(str1[0])
	month, err = strconv.Atoi(str1[1])
	day, err = strconv.Atoi(str1[2])
	if err != nil {
		return "", err
	}
	currentYear, _ := strconv.Atoi(time.Now().Format("2006"))
	if year > currentYear {
		return "", errors.New("年月日を正しく指定してください。")
	}
	if month < 1 || 12 < month {
		return "", errors.New("年月日を正しく指定してください。")
	}
	if day < 1 || 31 < day {
		return "", errors.New("年月日を正しく指定してください。")
	}

	// 時分秒の処理
	str2 := strings.Split(date[1], ":")
	if len(str2) != 3 {
		return "", errors.New("時分秒はHH:ii:ssで指定してください。")
	}
	hour, err = strconv.Atoi(str2[0])
	minute, err = strconv.Atoi(str2[1])
	second, err = strconv.Atoi(str2[2])
	if err != nil {
		return "", err
	}
	if hour < 0 || 23 < hour {
		return "", errors.New("時分秒を正しく指定してください。")
	}
	if minute < 0 || 59 < minute {
		return "", errors.New("時分秒を正しく指定してください。")
	}
	if second < 0 || 59 < second {
		return "", errors.New("時分秒を正しく指定してください。")
	}

	convertTime := time.Date(year, time.Month(month), day,
		hour, minute, second, 0, time.Local)
	convertTime = convertTime.Add(time.Duration(1) * time.Second)

	return string(convertTime.Unix()), nil
}

func getJson(url string) (string, error) {
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
