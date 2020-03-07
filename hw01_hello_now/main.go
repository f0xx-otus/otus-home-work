package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	fmt.Println("current time: ", time.Now())
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("Failed to get time, exit", err)
	}
	fmt.Println("exact time: ", ntpTime)
}
