package tool

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"io"
	"log"
	"net/http"
	"os"
)

// DownloadFFMPEG FFMPEGをHPからダウンロードしたのち、展開する
func DownloadFFMPEG() (ffmpegPath string) {

	if f, err := os.Stat("ffmpeg-4.2.2-amd64-static/ffmpeg"); os.IsNotExist(err) || f.IsDir() {
		if err := os.RemoveAll("ffmpeg-4.2.2-amd64-static"); err != nil {
			fmt.Println(err)
		}
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

		if err := os.RemoveAll("hoge"); err != nil {
			fmt.Println(err)
		}

		ffmpegPath = UnpackFFMPEG("ffmpeg-release-amd64-static.tar.xz")
		fmt.Println("FFMPEG ready.")
	} else {
		fmt.Println("Already get ffmpeg")
	}

	return
}

//unpackFFMPEG HPからダウンロードしたFFMPEGの圧縮ファイルをカレントディレクトリに展開
func UnpackFFMPEG(ffmpegPack string) (Path string) {

	tarxz := archiver.NewTarXz()
	err := tarxz.Unarchive(ffmpegPack, ".")
	if err != nil {
		log.Fatal(err)
	}
	return
}
