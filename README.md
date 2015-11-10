sana\_api\_client\_sample
=========================

Sana (紗凪) を叩くサンプルプログラムです.

詳細は, 各ディレクトリを見てください.

golang
------

大幅に変更しました。ちゃんとしたものが完成したらまた書きます。

現状では、`./soraClient account1 account2 ...`のように実行すると、
/anime/v1/twitter/follower/statusにリクエストを送り、最新のフォロワー数を
取得します。

その後、/anime/v1/twitter/follower/historyにリクエストを送ります。これの
リクエストパラメータは、

1. accountのみ指定した場合
2. accountとend_dateを指定した場合

の2種類あるので、それぞれを実行します。

### Compiling & Executing

#### Compile

```
$ go build sanaClient.go
```
