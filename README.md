# git-get

> Clone into a GOPATH like directory structure

## Install
```
go install github.com/icholy/git-get@latest
```

## Usage

```
git get [<options>] [--] <repo>
```

* Arguments are forwarded directly to `git clone`.
* The `<repo>` parameter **MUST** be placed last.

## Config

By default, repositories are cloned to `~/src`.
This can be changed with the `GIT_GET_PATH` env variable.

## Demo

![](tty.gif)
