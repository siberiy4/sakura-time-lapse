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

	for _, record := range event.Records {
		util.MakeDirectoriy("/tmp/sakura")
		util.MakeDirectoriy("/tmp/sakura/pre")
		util.MakeDirectoriy("/tmp/sakura/movie")
		util.MakeDirectoriy("/tmp/sakura/takumi/")
		util.MakeDirectoriy("/tmp/sakura/jpg")
		tool.DownloadFFMPEG()
		bucketName := record.S3.Bucket.Name
		fmt.Println(bucketName)
		filePath := record.S3.Object.Key
		fmt.Println(filePath)
		movieName := strings.Split(filePath, "/")[2]
		fmt.Println(movieName)
		timelapse.MakeTimeLapse(filePath, movieName, bucketName)
	}

}

func main() {
	fmt.Println("start")
	lambda.Start(ex)
}
