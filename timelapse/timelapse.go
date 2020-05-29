package timelapse

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"sakura-time-lapse/s3"
	"sakura-time-lapse/util"
)

// unpackTar jpgがまとめられたtarを展開する
func unpackTar(tarFile string) {
	var file *os.File
	var err error

	//tarのopen
	if file, err = os.Open(tarFile); err != nil {
		fmt.Println(err)
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
			fmt.Println(err)
		}

		buf := new(bytes.Buffer)
		if _, err = io.Copy(buf, reader); err != nil {
			fmt.Println(err)
		}
		if err = ioutil.WriteFile(fmt.Sprintf("/tmp/sakura/jpg/source%04d.jpg", i), buf.Bytes(), 0755); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("unpack tar")
}

//作ったmp４を連結してその日の
func unitMP4(movieName,bucketName string) {
	dest := "/tmp/sakura/movie/" + movieName + ".mp4"
	check:=s3.CheckObject("movie/" + movieName + ".mp4",bucketName)
	if  !check {
		if err := os.Rename("/tmp/sakura/pre/addition.mp4", dest); err != nil {
			fmt.Println(err)
		}
		fmt.Println("make mp4")

	} else {
		makeUnitTXT()
		s3.GetS3file("movie/"+movieName+ ".mp4","/tmp/sakura/pre/time-lapse.mp4",bucketName)
		err := exec.Command("/tmp/sakura/ffmpeg-4.2.1-amd64-static/ffmpeg", "-f", "concat", "-i", "/tmp/sakura/unitMP4.txt", "-c", "copy", dest, "-y").Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("unit mp4")
	}
	s3.UpMovie(dest,"movie/" + movieName + ".mp4",bucketName)

}

// MakeTimeLapse タイムラプスを作成する
func MakeTimeLapse(tarFile, movieName,bucketName string) {
	s3.GetS3file(tarFile, "/tmp/sakura/takumi/jpg.tar", bucketName)
	unpackTar("/tmp/sakura/takumi/jpg.tar")
	//タイムラプスの作成
	err := exec.Command("/tmp/sakura/ffmpeg-4.2.1-amd64-static/ffmpeg", "-f", "image2", "-r", "20", "-i", "/tmp/sakura/jpg/source%04d.jpg", "-r", "40", "-an", "-vcodec", "libx264", "-pix_fmt", "yuv420p", "/tmp/sakura/pre/addition.mp4", "-y").Run()
	if err != nil {
		fmt.Println("faild make time lapse")
		fmt.Println(err)
	}
	util.RemoveAllFile("/tmp/sakura/jpg/")
	unitMP4(movieName,bucketName)

	fmt.Println("make time-lapse")
}

// MakeUbitTXT aa
func makeUnitTXT() {
	text := fmt.Sprint(`file '/tmp/sakura/pre/time-lapse.mp4'
file '/tmp/sakura/pre/addition.mp4'`)
	//os.O_RDWRを渡しているので、同時に読み込みも可能
	file, err := os.OpenFile("/tmp/sakura/unitMP4.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//エラー処理
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Fprintln(file, text)

	fmt.Print(text)
}
