sana\_api\_client\_sample
=========================

Sana (紗凪) を叩くサンプルプログラムです.

詳細は, 各ディレクトリを見てください.

golang
------

### Compiling & Executing

#### Usage

```
$ ./sanaClient -a account1,account2,account3...
```

データを取得したい Twitter アカウントを-a オプションで指定します. 取得したい
アカウントが複数ある場合は, カンマ (,) 区切りでプログラムに渡します.

#### Compile

```
$ go build sanaClient.go
```

### Example

```
$ ./sanaClient -a usagi_anime,kinmosa_anime,aldnoahzero
{"aldnoahzero":{"updated_at":1447057264,"follower":47986},"usagi_anime":{"updated_at":1447057806,"follower":134747},"kinmosa_anime":{"updated_at":1447057445,"follower":68386}}
```
