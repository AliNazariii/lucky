package time

import (
	"fmt"
	"time"
)

func PrintTicker(duration time.Duration) {
	ticker := time.NewTicker(duration)
	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t.Format("15:04:05"))
		}
	}
}
