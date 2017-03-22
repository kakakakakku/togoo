# Togoo [![Build Status](https://travis-ci.org/kakakakakku/togoo.svg)](https://travis-ci.org/kakakakakku/togoo) [![Go Report Card](https://goreportcard.com/badge/github.com/kakakakakku/togoo)](https://goreportcard.com/report/github.com/kakakakakku/togoo)

Togoo : Simple CLI TODO Tool, using Golang and SQLite.

![Usage](images/usage.gif)

## Install

To install, use `go get`:

```bash
$ go get -u github.com/kakakakakku/togoo
$ togoo init
```

## Commands

```
init
add, a
update, u
list, l
done, d
delete
help, h
```

## Usage

```bash
➜  ~  togoo add xxx
➜  ~  togoo add yyy
➜  ~  togoo list
+----+-------+--------+
| NO | TITLE | STATUS |
+----+-------+--------+
|  1 | xxx   | -      |
|  2 | yyy   | -      |
+----+-------+--------+
➜  ~  togoo done 1
➜  ~  togoo list
+----+-------+--------+
| NO | TITLE | STATUS |
+----+-------+--------+
|  2 | yyy   | -      |
+----+-------+--------+
➜  ~  togoo list -a
+----+-------+--------+
| NO | TITLE | STATUS |
+----+-------+--------+
|  1 | xxx   | Done   |
|  2 | yyy   | -      |
+----+-------+--------+
```

## Contribution

1. Fork ([https://github.com/kakakakakku/togoo/fork](https://github.com/kakakakakku/togoo/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

* [kakakakakku](https://github.com/kakakakakku)

