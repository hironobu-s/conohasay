# conohasay

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE) [![Build Status](https://travis-ci.org/hironobu-s/conohasay.svg?branch=master)](https://travis-ci.org/hironobu-s/conohasay) [![codebeat badge](https://codebeat.co/badges/792c6579-ec06-4841-a6e2-d49df29c0640)](https://codebeat.co/projects/github-com-hironobu-s-conohasay)


**conohasay** は[cowsay](https://en.wikipedia.org/wiki/Cowsay)ライクなCLIツールです。ターミナル上で[ConoHa](https://www.conoha.jp/)のキャラクターが楽しくおしゃべりしますよ。


![conohasay](images/screen2.png)


## Install

以下の手順で実行ファイルをダウンロードしてください。以下のコマンドはカレントディレクトリにダウンロードしますが、使用頻度が高い場合はパスの通った場所に置いてください。

**Mac OSX**

```bash
curl -sL https://github.com/hironobu-s/conohasay/releases/download/current/conohasay-osx.amd64.gz | zcat > conohasay && chmod +x ./conohasay
```

**Linux(amd64)**

```bash
curl -sL https://github.com/hironobu-s/conohasay/releases/download/current/conohasay-linux.amd64.gz | zcat > conohasay && chmod +x ./conohasay
```

**Windows(amd64)**

[ZIP file](https://github.com/hironobu-s/conohasay/releases/download/current/conohasay.amd64.zip)

## 使い方

引数を渡すと、その文字列を表示します(標準入力から渡すこともできます)

```shell
conohasay I'm ConoHa Mikumo! 
```

![conohasay](images/screen1.png)


``-c`` オプションでキャラクターを変更できます。


```shell
conohasay -c anzu
```

![conohasay](images/screen3.png)


``-l`` オプションで``-c``オプションに指定できる値の一覧が取得できます。

```shell
# conohasay -l
anzu
conoha
logo
umemiya
```

``-s`` オプションでキャラクターのサイズを変更できます。指定できる値は``s``, ``m``, ``l``のどれかです。

```shell
conohasay -c logo -s l
```


その他のオプションは下記です。

* ``-W wrapcolumn`` メッセージの幅を指定します。デフォルトは50です。cowsayの同名オプションと同じ機能です。
* ``-f force-horizontal-layout`` キャラクターとメッセージを縦に並べることを強制します。デフォルトは端末の幅に応じて自動調整されます。

* ``-v`` バージョンを表示します
* ``-h`` ヘルプを表示します

## 活用例

fortuneコマンドの出力を渡すのがcowsayの定番ですが、**hostname**や**w**や**uptime**などコマンド出力を渡すのも面白いです。

```shell
w -s | conohasay -c conoha -s m
```

![conohasay](images/screen5.png)


## ライセンス

MIT
