//
// Copyright (c) 2015 Rirei Kuroshima <rireikuroshima@icloud.com>
//
package main

import (
	"fmt"
	"log"
	"./sana"
)

// sanaパッケージの使用例です。sanaから取得したJSONを標準出力に表示します。
func main() {
	// GetLatestFollower()関数を使って、スライスで指定したアニメ公式アカウント
	// の最新のフォロー数を取得する。
	accounts := []string{"usagi_anime", "kinmosa_anime", "aldnoahzero"}
	if json, err := sana.GetLatestFollower(accounts); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(json)
	}

	// GetFollowerHistory()関数を使って指定したアニメ公式アカウントの
	// フォロワー数の履歴を取得する。
	if json, err := sana.GetFollowerHistory("usagi_anime"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(json)
	}

	// GetFollowerHisotry()関数を使って指定したアニメ公式のアカウントの
	// フォロワー数を取得します。なお、第2引数で指定した日時より過去のデータを
	// を取得します。
	json, err := sana.GetFollowerHistory("usagi_anime",
		"2015-11-10 21:04:00")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(json)
	}
}
