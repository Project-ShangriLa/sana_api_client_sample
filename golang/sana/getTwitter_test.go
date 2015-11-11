package sana

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

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
