package io

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// File copies a single file from src to dst
func CopyFile(src, dst string) error {
	var err error
	var srcFile *os.File
	var dstFile *os.File
	var srcInfo os.FileInfo

	if srcFile, err = os.Open(src); err != nil {
		return err
	}
	defer srcFile.Close()

	if dstFile, err = os.Create(dst); err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return err
	}
	if srcInfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcInfo.Mode())
}

// Dir copies a whole directory recursively
func CopyDir(src string, dst string) error {
	var err error
	var fileInfos []os.FileInfo
	var srcInfo os.FileInfo

	if srcInfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	if fileInfos, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fileInfo := range fileInfos {
		srcFile := path.Join(src, fileInfo.Name())
		dstFile := path.Join(dst, fileInfo.Name())

		if fileInfo.IsDir() {
			if err = CopyDir(srcFile, dstFile); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = CopyFile(srcFile, dstFile); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
