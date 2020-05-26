package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	time.Local = time.FixedZone("JST", +9*60*60)
	ticker := time.NewTicker(20 * time.Second)
	lastMinute := time.Now().Minute()
	defer ticker.Stop()
	for {
		select {
		case now := <-ticker.C:
			fmt.Println(now)
			if (now.Minute()%10 == 2 && (5 <= now.Hour() && now.Hour() < 20 || now.Hour() == 4 && now.Minute()/10 > 0) || now.Hour() == 20 && now.Minute() == 2) && now.Minute() != lastMinute {
				path := tarPath(now)
				if _, err := os.Stat("/mnt/sakura/"+path); os.IsNotExist(err) {
				}else{
					copyFile("/mnt/sakura/"+path,"mnt/test/"+path)
				}
				lastMinute = now.Minute()
			}
		}
	}
}

func copyFile(s string, d string) {
	w, err := os.Create(d)
	if err != nil {
		fmt.Println(err)
	}

	r, err := os.Open(s)
	if err != nil {
		fmt.Println(err)
	}

	_, err = io.Copy(w, r)
	if err != nil {
		fmt.Println(err)
	}
}

func tarPath(t time.Time) (path string) {
	min := 0
	hour := 0
	if t.Minute() < 10 {
		min = 5
		hour = t.Hour() - 1
	} else {
		min = (t.Minute() / 10) - 1
		hour = t.Hour()
	}
	return fmt.Sprintf("takumi/jpg/%d/%02d%02d%02d/%02d%1d.tar", t.Year(), t.Year()%100, t.Month(), t.Day(), hour, min)

}
