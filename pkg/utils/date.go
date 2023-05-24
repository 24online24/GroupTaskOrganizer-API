package utils

import (
	"fmt"
	"time"
)

func DateGen(dayDiff int) string {
	currentTime := time.Now()
	return fmt.Sprintf("%d-%02d-%02d", currentTime.Year(), currentTime.Month(), currentTime.Day()+dayDiff)
}
