package main

import (
	"fmt"
	"os"

	"github.com/Armatorix/CronParser/pkg/cron"
)

func main() {
	cron, err := cron.NewFromOsArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Execution failed: ", err)
		os.Exit(-1)
	}
	fmt.Println(cron)
}
