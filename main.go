package main

import (
	"context"
	"fmt"
	"sakura-time-lapse/timelapse"
	"sakura-time-lapse/tool"
	"sakura-time-lapse/util"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func ex(ctx context.Context, event events.S3Event) {
	util.MakeDirectoriy("/tmp/sakura")
	util.MakeDirectoriy("/tmp/sakura/pre")
	util.MakeDirectoriy("/tmp/sakura/movie")
	util.MakeDirectoriy("/tmp/sakura/takumi/")
	util.MakeDirectoriy("/tmp/sakura/takumi/jpg")
	util.MakeDirectoriy("/tmp/sakura/jpg")
	tool.DownloadFFMPEG()

	for _, record := range event.Records {
		bucketName := record.S3.Bucket.Name
		filePath := record.S3.Object.Key
		movieName:=strings.Split(filePath,"/")[2]
		timelapse.MakeTimeLapse(filePath,movieName,bucketName)
	}

}

func main() {
	fmt.Println("start")
	lambda.Start(ex)
}
