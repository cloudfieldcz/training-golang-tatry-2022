package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello world")
	log.SetLevel(log.InfoLevel)
	log.Info("Starting app version: ", time.Now())
}
