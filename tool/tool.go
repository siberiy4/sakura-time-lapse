package tool

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mholt/archiver/v3"
)

// DownloadFFMPEG FFMPEGをHPからダウンロードしたのち、展開する
func DownloadFFMPEG() (ffmpegPath string) {

	if f, err := os.Stat("/tmp/sakura/ffmpeg-4.2.2-amd64-static/ffmpeg"); os.IsNotExist(err) || f.IsDir() {
		if err := os.RemoveAll("/tmp/sakura/ffmpeg-4.2.2-amd64-static"); err != nil {
			fmt.Println(err)
		}
		ffmpegURL := "https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz"

		resp, err := http.Get(ffmpegURL)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		out, err := os.Create("/tmp/sakura/ffmpeg-release-amd64-static.tar.xz")
		if err != nil {
			fmt.Println(err)
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		ffmpegPath = unpackFFMPEG("/tmp/sakura/ffmpeg-release-amd64-static.tar.xz")
		if err := os.Remove("/tmp/sakura/ffmpeg-release-amd64-static.tar.xz"); err != nil {
			fmt.Println(err)
		}
		fmt.Println("FFMPEG ready.")
	} else {
		fmt.Println("Already get ffmpeg")
	}

	return
}

//unpackFFMPEG HPからダウンロードしたFFMPEGの圧縮ファイルをカレントディレクトリに展開
func unpackFFMPEG(ffmpegPack string) (Path string) {

	tarxz := archiver.NewTarXz()
	err := tarxz.Unarchive(ffmpegPack, "/tmp/sakura/")
	if err != nil {
		fmt.Println(err)
	}
	return
}
