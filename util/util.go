package util

import (
	"fmt"
	"io"
	"os"
)

// CopyFile sのファイルをdとしてコピー
func CopyFile(s string, d string) {

	w, err := os.Create(d)
	defer w.Close()
	if err != nil {
		fmt.Println(err)
	}

	r, err := os.Open(s)
	defer r.Close()
	if err != nil {
		fmt.Println(err)
	}

	_, err = io.Copy(w, r)
	if err != nil {
		fmt.Println(err)
	}
}

// MakeDirectoriy nameを作成
func MakeDirectoriy(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		os.Mkdir(name, 0777)
	}
}

// RemoveAllFile 引数のディレクトリ内のファイルをすべて消す
func RemoveAllFile(directoryPath string) (err error) {
	if err := os.RemoveAll(directoryPath); err != nil {
		return err
	}

	if err := os.Mkdir(directoryPath, 0777); err != nil {
		return err
	}

	return nil
}
