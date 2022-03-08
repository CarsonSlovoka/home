/*
coverage 81.8%
*/

package io_test

import (
	io2 "carson.io/pkg/io"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

func TestCopyFileAndCopyDir(t *testing.T) {
	var err error
	{ // Create test dir
		for _, dirPath := range []string{
			"./temp/css/",
			"./temp/js/",
			"./temp/js/src",
		} {
			if err = os.MkdirAll(dirPath, fs.ModeType); err != nil {
				t.Fatalf(err.Error())
			}
		}
	}

	{ // Create test files
		for _, arg := range []struct {
			filePath string
			content  string
		}{
			{"./temp/css/a.css", "hello world"},
			{"./temp/js/main.js", "main.js"},
			{"./temp/js/src/b.js", "b.js"},
		} {
			var file *os.File
			if file, err = os.Create(arg.filePath); err != nil {
				t.Fatalf(err.Error())
			}
			_, _ = file.WriteString(arg.content)
			_ = file.Close()
		}
	}

	{ // Test Copy File
		if err = io2.CopyFile("temp/css/a.css", "a_copy.css"); err != nil {
			t.FailNow()
		}

		var fileBytes []byte
		if fileBytes, err = os.ReadFile("a_copy.css"); err != nil {
			t.FailNow()
		}
		if string(fileBytes) != "hello world" {
			t.FailNow()
		}
	}

	{ // Test Copy Dir
		if err = io2.CopyDir("temp", "temp_dst"); err != nil {
			t.FailNow()
		}
		for _, arg := range []struct {
			filePath string
			content  string
		}{
			{"./temp_dst/css/a.css", "hello world"},
			{"./temp_dst/js/main.js", "main.js"},
			{"./temp_dst/js/src/b.js", "b.js"},
		} {
			var fileBytes []byte
			if fileBytes, err = os.ReadFile(arg.filePath); err != nil {
				t.Fatalf(err.Error())
			}
			if string(fileBytes) != arg.content {
				t.FailNow()
			}
		}
	}

	{ // delete test dir
		for _, testDir := range []string{
			"./temp",
			"./temp_dst",
			"a_copy.css",
		} {
			if err = os.RemoveAll(testDir); err != nil {
				t.FailNow()
			}
		}
	}
}

func TestCopyErr(t *testing.T) {
	var err error
	if err = io2.CopyFile("notExist.txt", "dst.txt"); err == nil {
		t.FailNow()
	}
	fmt.Println(err.Error()) // The system cannot find the file specified.

	if err = io2.CopyDir("notExistDir", "dst"); err == nil {
		t.FailNow()
	}
	fmt.Println(err.Error()) // The system cannot find the file specified.

	if err = io2.CopyFile("copy.go", "NotExistDir/foo.txt"); err == nil {
		t.FailNow()
	}
	fmt.Println(err.Error()) // The system cannot find the path specified.
}
