package s3

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// GetS3file 引数で指定したfile pathでS3からファイルを取得
func GetS3file(s3File, filename, bucket string) {

	sess := session.Must(session.NewSession())
	
	
	// Create a file to write the S3 Object contents to.
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to create file %q, %v", filename, err))
		return
	}
	// Create a downloader with the session and default options
	downloader:=s3manager.NewDownloader(sess)
	// Write the contents of S3 Object to the file

	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(s3File),
	})
	if err != nil {
		fmt.Println(fmt.Errorf("failed to download file, %v", err))
		return
	}

}

//UpMovie S3にファイルをアップロード
func UpMovie(filename,s3File,  bucket string)  {
	sess := session.Must(session.NewSession())
	
	
	uploader:=s3manager.NewUploader(sess)

	file,err:=os.Open(filename)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer file.Close()

	_,err=uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(s3File),
		Body: file,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}