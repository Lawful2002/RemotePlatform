package cron

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func removeOldContainers() error {
	fmt.Println("Remove Container Function")
	return nil
}

func RunCronJob() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minutes().Do(func() {
		removeOldContainers()
	})
	s.StartBlocking()
}
