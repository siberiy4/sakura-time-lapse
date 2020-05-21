package scheduler

import (
	"fmt"
	"sakura-time-lapse/timelapse"
	"time"
)

// Scheduler 一定時間ごとに実行する
func Scheduler() {
	ticker := time.NewTicker(10 * time.Second)
	lastMinute := time.Now().Minute()
	defer ticker.Stop()
	for {
		select {
		case now := <-ticker.C:
			fmt.Println(now.Format(time.RFC3339))
			if now.Minute()%10 == 1 && now.Minute() != lastMinute {
				timelapse.MakeTimeLapse()
				lastMinute = now.Minute()
			}
		}
	}
}
