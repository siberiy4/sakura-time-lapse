package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Getjpgtar 引数で指定したfile pathでS3からタイムラプス用の画像を取得
func Getjpgtar(filePath string,filename string) {
	//ACCESS_KEYとSECRET_KEYを.envから読む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	creds := credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY"), os.Getenv("SECRET_KEY"), "")

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: creds,
		Region: aws.String(os.Getenv("S3_REGION")),
	}))

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	// Create a file to write the S3 Object contents to.
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to create file %q, %v", filename, err))
		return
	}

	// Write the contents of S3 Object to the file
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(filePath),
	})
	if err != nil {
		fmt.Println(fmt.Errorf("failed to download file, %v", err))
		return
	}
	fmt.Printf("jpg-file downloaded, %d bytes\n", n)

}
