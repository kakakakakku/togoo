# Togoo

[![Build Status](https://travis-ci.org/Kakakakakku/togoo.svg)](https://travis-ci.org/Kakakakakku/togoo)

## Description

Simple CLI TODO tool, using Golang and SQLite.

## Install

To install, use `go get`:

```bash
$ go get -u github.com/kakakakakku/togoo
$ togoo init
```

## Commands

```
init
add
list
done
delete
help, h
```

## Usage

```bash
$ togoo add sample_task1
$ togoo add sample_task2
$ togoo list
+----+--------------+--------+
| NO |    TITLE     | STATUS |
+----+--------------+--------+
|  1 | sample_task1 | -      |
|  2 | sample_task2 | -      |
+----+--------------+--------+
$ togoo done 1
$ togoo list
+----+--------------+--------+
| NO |    TITLE     | STATUS |
+----+--------------+--------+
|  1 | sample_task1 | Done   |
|  2 | sample_task2 | -      |
+----+--------------+--------+
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
