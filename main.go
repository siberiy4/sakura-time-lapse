package main

import (
	"fmt"
	"sakura-time-lapse/tool"
	"sakura-time-lapse/util"
	"sakura-time-lapse/scheduler"
)

func main() {

	fmt.Println("start")
	util.MakeDirectoriy("/tmp/sakura")
	util.MakeDirectoriy("/tmp/sakura/pre")
	util.MakeDirectoriy("/tmp/sakura/movie")
	util.MakeDirectoriy("/tmp/sakura/jpg")
	tool.DownloadFFMPEG()
	scheduler.Scheduler()
}
