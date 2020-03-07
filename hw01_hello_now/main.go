package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	fmt.Println("current time: ", time.Now())
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println("Failed to get time, exit", err)
		os.Exit(1)
	}
	fmt.Println("exact time: ", ntpTime)
}
