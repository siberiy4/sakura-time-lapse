package util

import (
	"io"
	"log"
	"os"
)

// CopyFile sのファイルをdとしてコピー
func CopyFile(s string, d string) {
	w, err := os.Create(d)
	if err != nil {
		log.Fatal(err)
	}

	r, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(w, r)
	if err != nil {
		log.Fatal(err)
	}
}

// MakeDirectoriy nameを作成
func MakeDirectoriy(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		os.Mkdir(name, 0777)
	}
}
