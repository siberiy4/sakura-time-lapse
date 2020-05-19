package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
	"github.com/mholt/archiver/v3"
)

//tar の展開
func unpackTar() {
	var file *os.File
	var err error

	//tarのopen
	if file, err = os.Open("test/200506-125.tar"); err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	//tar にまとめられているfileの先頭のfileを受け取る
	reader := tar.NewReader(file)

	//一つずつfileを作っていく
	for i := 0; true; i++ {
		_, err = reader.Next()
		if err == io.EOF {
			// ファイルの最後
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		buf := new(bytes.Buffer)
		if _, err = io.Copy(buf, reader); err != nil {
			log.Fatalln(err)
		}
		//fileの作成
		if err = ioutil.WriteFile(fmt.Sprintf("test/source%04d.jpg", i), buf.Bytes(), 0755); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("unpack tar")
}

func makeTimeLapse() {
	unpackTar()
	//タイムラプスの作成
	err := exec.Command("./ffmpeg-4.2.2-amd64-static/ffmpeg", "-f", "image2", "-r", "20", "-i",  "test/source%04d.jpg", "-r", "40", "-an", "-vcodec", "libx264",  "-pix_fmt", "yuv420p", "video.mp4", "-y").Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("make time-lapse")
}

func scheduler() {
	ticker := time.NewTicker(30 * time.Second)
	lastMinute := time.Now().Minute()
	defer ticker.Stop()
	for {
		select {
		case now := <-ticker.C:
			fmt.Println(now.Format(time.RFC3339))
			if now.Minute()%10 == 1 && now.Minute() != lastMinute {
				makeTimeLapse()
				lastMinute = now.Minute()
			}
		}
	}
}

func makeDirectories() {
	if _, err := os.Stat("./test"); os.IsNotExist(err) {
		os.Mkdir("./test", 0777)
	}
	if _, err := os.Stat("./movie"); os.IsNotExist(err) {
		os.Mkdir("./movie", 0777)
	}


}

func downloadFFMPEG() {
	ffmpegURL := "https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz"

	resp, err := http.Get(ffmpegURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	out, err := os.Create("ffmpeg-release-amd64-static.tar.xz")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	unpackFFMPEG()
}

func unpackFFMPEG() {
	tarxz:=archiver.NewTarXz()
	err:=tarxz.Unarchive("ffmpeg-release-amd64-static.tar.xz",".")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {/*
	makeDirectories()
	downloadFFMPEG()
	scheduler()*/
	makeTimeLapse()
}
