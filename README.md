sana\_api\_client\_sample
=========================

Sana (紗凪) を叩くサンプルプログラムです.

詳細は, 各ディレクトリを見てください.

golang
------

### Compiling & Executing

#### Usage

```
sanaClient account1 account2 account3 ...
```

データを取得したい Twitter アカウントをコマンドラインで指定します.

#### Compile

```
$ go build sanaClient.go
```

### Example

```
$ ./sanaClient usagi_anime kinmosa_anime aldnoahzero
{"aldnoahzero":{"updated_at":1447057264,"follower":47986},"usagi_anime":{"updated_at":1447057806,"follower":134747},"kinmosa_anime":{"updated_at":1447057445,"follower":68386}}
```
