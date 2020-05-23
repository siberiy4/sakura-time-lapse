package main

import (
	"fmt"
	"sakura-time-lapse/tool"
	"sakura-time-lapse/util"
)

func main() {

	fmt.Println("start")
	util.MakeDirectoriy("pre")
	util.MakeDirectoriy("movie")
	util.MakeDirectoriy("jpg")
	tool.DownloadFFMPEG()
}
