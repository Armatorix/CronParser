[![codecov](https://codecov.io/gh/Armatorix/CronParser/branch/main/graph/badge.svg?token=IVZ5PJTZBF)](https://codecov.io/gh/Armatorix/CronParser)

# CronParser

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
