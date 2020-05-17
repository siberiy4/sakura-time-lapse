package main

import (
	"archive/tar"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"log"
	"fmt"
)

func main() {
	var file *os.File
	var err error

	if file, err = os.Open("test/200506-125.tar"); err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// ReaderはClose()はない。
	reader := tar.NewReader(file)

	for i:=0;true;i++ {
		_, err = reader.Next()
		if err == io.EOF {
			// ファイルの最後
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		buf := new(bytes.Buffer)
		if _, err = io.Copy(buf, reader); err != nil {
			log.Fatalln(err)
		}

		if err = ioutil.WriteFile( fmt.Sprintf("test/source%4d.jpg",i)  , buf.Bytes(), 0755); err != nil {
			log.Fatal(err)
		}
	}
}