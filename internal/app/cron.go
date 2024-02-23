package app

import (
	"fmt"
	"time"
)

func startCron() {
	now := time.Now()
	firstTime := time.Date(now.Year(), now.Month(), now.Day(), 16, 05, 0, 0, now.Location())

	if now.After(firstTime) {
		firstTime = firstTime.Add(24 * time.Hour)
	}

	delay := firstTime.Sub(now)
	timer := time.NewTimer(delay)

	task := func() {
		fmt.Println("task is processing")
	}

	go func() {
		for {
			<-timer.C
			task()
			timer.Reset(24 * time.Hour)
		}
	}()

	select {}
}
