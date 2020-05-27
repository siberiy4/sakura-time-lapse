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
				path,dir,mat := tarPath(now)
				makeDirectoriy(dir)
				if _, err := os.Stat("/mnt/sakura/"+path); os.IsNotExist(err) {
				}else{
					copyFile("/mnt/sakura/"+path,"/mnt/s3/"+mat)
				}
				lastMinute = now.Minute()
			}
		}
	}
}

func copyFile(s string, d string) {
	w, err := os.Create(d)
	defer w.Close()
	if err != nil {
		fmt.Println(err)
	}

	r, err := os.Open(s)
	defer r.Close()
	if err != nil {
		fmt.Println(err)
	}

	_, err = io.Copy(w, r)
	if err != nil {
		fmt.Println(err)
	}
}

func makeDirectoriy(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		os.Mkdir(name, 0777)
	}
}

func tarPath(t time.Time) (lcPath,testdirectory,testFile string) {
	min := 0
	hour := 0
	if t.Minute() < 10 {
		min = 5
		hour = t.Hour() - 1
	} else {
		min = (t.Minute() / 10) - 1
		hour = t.Hour()
	}
	return fmt.Sprintf("takumi/jpg/%d/%02d%02d%02d/%02d%1d.tar", t.Year(), t.Year()%100, t.Month(), t.Day(), hour, min),fmt.Sprintf("takumi/jpg/%02d%02d%02d/", t.Year()%100, t.Month(), t.Day()),fmt.Sprintf("takumi/jpg/%02d%02d%02d/%02d%1d.tar", t.Year()%100, t.Month(), t.Day(), hour, min)

}
