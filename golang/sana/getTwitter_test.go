// sanaパッケージの各関数のテスト。
//     o getJson()関数のテストの仕方がいまいち分からない
//     o changeTest()関数はテストケースがまだ少ない
//     o GetLatestFollower()とGetFollowerHistory()関数は、実行するたびにJSONの
//       データが変わるので、どうやったら良いか考えるのがめんどくさい
//     o GetFollowerHistory()関数は引数が可変で、それをどうやったらうまいこと
//       テストが書けるのかが分からない
// 要するに、このテストはまだまだ。
package sana

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// 引数にURL（sanaのURL）を指定したら、そのJSONデータを取得する関数である、
// getJson()関数のテスト関数です。
func TestGetJson(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "test")
		}))
	defer ts.Close()

	defaultProxy := http.DefaultTransport.(*http.Transport).Proxy
	http.DefaultTransport.(*http.Transport).Proxy =
		func(req *http.Request) (*url.URL, error) {
		// テストサーバのURLをProxyに設定
		return url.Parse(ts.URL)
	}
	// テストが終わったらProxyを元に戻す
	defer func() {
		http.DefaultTransport.(*http.Transport).Proxy = defaultProxy
	}()

	actual, _ := getJson("http://hoge.com")
	expected := "test"
	if actual != expected {
		t.Errorf("Expected %v, but %v ", expected, actual)
	}
}

// TestChangeTime()関数で使用するテストケース。
var testdata = []struct {
	in string
	out string
	e string
}{
	{"2015-11-05 15:04:05", "1446703445", ""},
	{"2015-11-05_15:04:05", "", "YYYY-MM-DD HH:ii:ssで指定してください。"},
	{"201511-05 15:04:05", "", "年月日はYYYY-MM-DDで指定してください。"},
	{"2015-1105 15:04:05", "", "年月日はYYYY-MM-DDで指定してください。"},
	{"20151105 15:04:05", "", "年月日はYYYY-MM-DDで指定してください。"},
	{"2015-11-05-06 15:04:05", "", "年月日はYYYY-MM-DDで指定してください。"},
	{"2016-11-05 15:04:05", "", "年月日を正しく指定してください。"},
	{"2015-13-05 15:04:05", "", "年月日を正しく指定してください。"},
	{"2015-00-05 15:04:05", "", "年月日を正しく指定してください。"},
	{"2015-11-32 15:04:05", "", "年月日を正しく指定してください。"},
	{"2015-11-00 15:04:05", "", "年月日を正しく指定してください。"},
	{"2015-11-05 1504:05", "", "時分秒はHH:ii:ssで指定してください。"},
	{"2015-11-05 15:0405", "", "時分秒はHH:ii:ssで指定してください。"},
	{"2015-11-05 150405", "", "時分秒はHH:ii:ssで指定してください。"},
	{"2015-11-05 24:04:05", "", "時分秒を正しく指定してください。"},
	{"2015-11-05 15:60:05", "", "時分秒を正しく指定してください。"},
	{"2015-11-05 15:04:60", "", "時分秒を正しく指定してください。"},
}

// YYYY-MM-DD HH:ii:ss形式で渡された日時をUnix時間に変換する関数である、
// changeTime()関数のテスト関数。
func TestChangeTime(t *testing.T) {
	for _, tt := range testdata {
		actual, actualErr := changeTime(tt.in)

		if tt.out != "" {
			expected := tt.out
			if actual != expected {
				t.Errorf("Expected: %v, but %v: ", expected,
					actual)
			}
		} else {
			expectedErr := errors.New(tt.e)
			if actualErr.Error() != expectedErr.Error() {
				t.Errorf("\n\tExpected: %v\n\tActual: %v",
					expectedErr, actualErr)
			}
		}
	}
}
