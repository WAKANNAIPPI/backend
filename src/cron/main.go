package main

import (
	"os/exec"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("@every 1h", req)
	c.Start()
	time.Sleep(2 * time.Minute)
}

func req() {
	exec.Command("curl", "localhost:8080/flag").Output()
}
