# CronParser

[![codecov](https://codecov.io/gh/Armatorix/CronParser/branch/main/graph/badge.svg?token=IVZ5PJTZBF)](https://codecov.io/gh/Armatorix/CronParser)
[![CircleCI](https://circleci.com/gh/Armatorix/CronParser/tree/main.svg?style=shield)](https://app.circleci.com/pipelines/github/Armatorix/CronParser)
[![Go Report Card](https://goreportcard.com/badge/github.com/Armatorix/CronParser)](https://goreportcard.com/report/github.com/Armatorix/CronParser)

## Requirements

- [go1.16](https://golang.org/doc/install)

## Installation

```bash
go get -u github.com/Armatorix/CronParser/cmd/cronparser
```

## Example run

```bash
cronparser "* * * * * /usr/bin/time"
```

## TODO

- tests for inapropreate execution
- handling for step asterisk
- move parsing/transformation functions to utility package
