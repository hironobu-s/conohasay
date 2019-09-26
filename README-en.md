# conohasay

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE) [![Build Status](https://travis-ci.org/hironobu-s/conohasay.svg?branch=master)](https://travis-ci.org/hironobu-s/conohasay) [![codebeat badge](https://codebeat.co/badges/792c6579-ec06-4841-a6e2-d49df29c0640)](https://codebeat.co/projects/github-com-hironobu-s-conohasay)

**conohasay** is a command-line program like [cowsay](https://en.wikipedia.org/wiki/Cowsay). It's generates the text message with ASCII picture of [ConoHa](https://www.conoha.jp/) family.

![conohasay](images/screen2.png)

## System requirements

We support Linux and macOS platform. **Your terminal should support 256 colors.**

## Install

### macOS

```bash
curl -sL https://github.com/hironobu-s/conohasay/releases/download/current/conohasay-osx.amd64.gz | zcat > conohasay && chmod +x ./conohasay
./conohasay -h
```

### Linux(amd64)

```bash
curl -sL https://github.com/hironobu-s/conohasay/releases/download/current/conohasay-linux.amd64.gz | zcat > conohasay && chmod +x ./conohasay
./conohasay -h
```
### Docker

If you have installed Docker, you will be able to try **conohasay** quickly.

```bash
docker run -i hironobu/conohasay -h
```

## How to use

Now that you've installed **conohasay**, you can make it work like so:

```shell
conohasay "I'm ConoHa Mikumo!"
```

![conohasay](images/screen1.png)

To change the character, use **-c** option.

```shell
conohasay -c anzu
```

![conohasay](images/screen3.png)

And you get the list of character by using **-l** option.


```shell
# conohasay -l
anzu
conoha
logo
umemiya
```

If you want to change the size of the picture, you may use **-s** option.

```shell
conohasay -c logo -s l
```

Other options:

* ``-W`` (wrapcolumn) Specifies roughly where the message should be wrapped
* ``-f`` (force-vertical-layout) Force the vertical layout (Default is auto-detected as the terminal width)
* ``-v`` Show version
* ``-h`` Show help

# License

MIT
