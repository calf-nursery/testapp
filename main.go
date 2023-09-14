package main

import (
	"fmt"

	"github.com/calf-nursery/testapp/log"
)

func main() {
	fmt.Println("Testapp - v0.0.5")
	fmt.Println("are you having fun yet")

	logger := log.New()
	logger.Info("testing info")
	logger.Debug("testing debug")
}
