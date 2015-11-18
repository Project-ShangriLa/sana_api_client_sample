sana\_api\_client\_sample
=========================

Sana (紗凪) を叩くサンプルプログラムです.

詳細は, 各ディレクトリを見てください.

golang
------

大幅に変更しました。ちゃんとしたものが完成したらまた書きます。

現状では、`./soraClient account1 account2 ...`のように実行すると、/anime/v1/twitter/follower/statusにリクエストを送り、最新のフォロワー数を取得します。

その後、/anime/v1/twitter/follower/historyにリクエストを送ります。これのリクエストパラメータは、

1. accountのみ指定した場合
2. accountとend\_dateを指定した場合

の2種類あるので、それぞれを実行します。

### Compiling & Executing

#### Compile

```
$ go build sanaClient.go
```

#### Executing

```
$ ./sanaClient account1 account2 account3 ...
```

引数で指定するアカウントは、`GetLatestFollower()`の引数として使われます。`GetFollowerHistory()`の実行は、

1. アカウントのみ指定する場合
2. アカウントと日時を指定する場合

の2パターン実行します。アカウントは、`usagi_anime`が指定されています。また、日時もプログラム内に固定しているので、変更したい人は、変更後コンパイルし直してください。

### Testing

`golang/sana`ディレクトリの中に`*_test.go`というファイルがあり、各関数のテストコードが書かれています。テスト内容はまだまだなので、改良する必要があります。以下、テストの結果です。

```
$ cd /path/to/golang/sana/
$ go test .
ok  	_/home/rirei/workspace/Project-ShangriLa/sana_api_client_sample/golang/sana	0.003s
```

また、テストの詳細を表示したい場合は、`-v`オプションを付けて実行します。

```
$ cd /path/to/golang/sana/
$ go test -v .
=== RUN   TestGetJson
--- PASS: TestGetJson (0.00s)
=== RUN   TestChangeTime
--- PASS: TestChangeTime (0.00s)
PASS
ok  	_/home/rirei/workspace/Project-ShangriLa/sana_api_client_sample/golang/sana	0.003s
```
