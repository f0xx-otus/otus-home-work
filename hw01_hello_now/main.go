package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	//hw02_unpack_string.Unpack("a4bc2d5e")
	fmt.Println("current time:", time.Now())
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalln("Failed to get time, exit", err)
	}
	fmt.Println("exact time:", ntpTime)
}
