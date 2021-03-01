package main

import (
	"fmt"

	"github.com/Armatorix/CronParser/pkg/cron"
)

func main() {
	fmt.Println(cron.NewFromOsArgs())
}
