package timelapse

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)


// unpackTar jpgがまとめられたtarを展開する
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
	err := exec.Command("./ffmpeg-4.2.2-amd64-static/ffmpeg", "-f", "image2", "-r", "20", "-i", "test/source%04d.jpg", "-r", "40", "-an", "-vcodec", "libx264", "-pix_fmt", "yuv420p", "video.mp4", "-y").Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("make time-lapse")
}

