package scheduler

import (
	"fmt"
	"sakura-time-lapse/timelapse"
	"sakura-time-lapse/s3"
	"time"
)

// Scheduler 一定時間ごとに実行する
func Scheduler() {
	time.Local = time.FixedZone("JST", +9*60*60)
	ticker := time.NewTicker(30 * time.Second)
	lastMinute := time.Now().Minute()
	defer ticker.Stop()
	for {
		select {
		case now := <-ticker.C:
			fmt.Println(now)
			if (now.Minute()%10 == 1&&5<=now.Hour()&&now.Hour()<8||now.Hour()==8&&now.Minute()==1) && now.Minute() != lastMinute {
				s3.Getjpgtar(s3.S3TarPath(now),"jpg/jpg.tar")
				timelapse.MakeTimeLapse()
				lastMinute = now.Minute()
			}
		}
	}
}
