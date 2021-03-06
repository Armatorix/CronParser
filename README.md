# CronParser

[![codecov](https://codecov.io/gh/Armatorix/CronParser/branch/main/graph/badge.svg?token=IVZ5PJTZBF)](https://codecov.io/gh/Armatorix/CronParser)
[![CircleCI](https://circleci.com/gh/Armatorix/CronParser/tree/main.svg?style=shield)](https://app.circleci.com/pipelines/github/Armatorix/CronParser)
[![Go Report Card](https://goreportcard.com/badge/github.com/Armatorix/CronParser)](https://goreportcard.com/report/github.com/Armatorix/CronParser)

## Requirements

- [go1.16](https://golang.org/doc/install)
- [make](https://man7.org/linux/man-pages/man1/make.1.html)

## Installation

```bash
go get -u github.com/Armatorix/CronParser/cmd/cronparser
```

**REMAMBER** to add the $GOPATH/bin to the $PATH environment variable

## Build from sources

```bash
make build
```

## Test

```bash
make test
```

## Example run

```bash
$ cronparser "37 21 */2 3-7,10-12 */2,3 /usr/bin/time"
minute         37
hour           21
day of month   1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31
month          3 4 5 6 7 10 11 12
day of week    0 2 3 4 6
command        /usr/bin/time

```

## TODO

- tests for inapropreate execution
- handling for step asterisk
- move parsing/transformation functions to utility package
