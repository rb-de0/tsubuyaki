#### 仕事中にターミナルからつぶやいちゃお的な何か

[![Build Status](https://travis-ci.org/rb-de0/tsubuyaki.svg?branch=master)](https://travis-ci.org/rb-de0/tsubuyaki)

## Golangのバージョン
以下の環境で開発・検証しました。

```bash
$ go version
go version go1.6.2 darwin/amd64
```

## インストール方法


#### リポジトリをダウンロード

```bash
go get github.com/rb-de0/tsubuyaki
```

#### 依存関係解決

リポジトリのディレクトリに移動して

```bash
gom install
```

#### ビルド

```bash
go build
```

#### Config.jsonの用意

ツイッターの認証に使う情報をConfig.jsonに記載してください。

#### 実行

```bash
./tsubuyaki tweet hoge
```

バイナリをパスの通っているディレクトリ配下に移動するか、リポジトリのディレクトリにパスを通すと良いかもしれません。

## 実装予定
きまぐれで作ったので実装するかどうか分かりませんが...

- [] ホームタイムライン表示
- [] リスト表示
- [] リツイート
- [] お気に入り
- [] リプライ
- [] 常時起動（API叩くたびにプロセス生成しないようにする）

## ライセンス
tsubuyakiは、MITライセンス下でリリースされています。 [MIT License](https://github.com/rb-de0/tsubuyaki/blob/master/LICENSE)

## 依存ライブラリ
依存ライブラリは、[DEPENDENCY.md](https://github.com/rb-de0/tsubuyaki/blob/master/DEPENDENCY.md) に記載しています。
