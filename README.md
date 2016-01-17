sana\_api\_client\_sample
=========================

Sana (紗凪) を叩くサンプルプログラムです.

詳細は, 各ディレクトリを見てください.

golang
------

### Compiling & Executing

#### Compile

```
$ cd golang/
$ go build example.go
```

#### Executing

```
$ ./example
```

sana/ディレクトリの中にあるgetTwitter.goの関数を一通り実行します。詳しくはソースを見てください。

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
