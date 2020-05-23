package main

import (
	"fmt"
	"sakura-time-lapse/tool"
	"sakura-time-lapse/util"
	"sakura-time-lapse/scheduler"
)

func main() {

	fmt.Println("start")
	util.MakeDirectoriy("pre")
	util.MakeDirectoriy("movie")
	util.MakeDirectoriy("jpg")
	tool.DownloadFFMPEG()
	scheduler.Scheduler()
}
