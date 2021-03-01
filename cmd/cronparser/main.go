package main

import (
	"fmt"

	"github.com/Armatorix/CronParser/pkg/cron"
)

func main() {
	cron, err := cron.NewFromOsArgs()
	if err != nil {
		fmt.Println("Execution failed: ", err)
	}
	fmt.Println(cron)
}
