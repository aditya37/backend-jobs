/*
 * File Name TemporyFile.go
 * Created on Sun Oct 04 2020
 *
 * Copyright (c) 2020
 */

package util

import (
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func CreateTemporyFile(FolderName,Pattern string, File multipart.File) (f *os.File,err error) {

	TempFile,err := ioutil.TempFile(FolderName,Pattern)
	if err != nil {
		return nil,err
	}
	defer TempFile.Close()
	CreateTemp,err := ioutil.ReadAll(File)
	if err != nil {
		return nil,err
	}
	TempFile.Write(CreateTemp)
	
	return TempFile,nil
}

func GetContentType(file multipart.File) (string,error) {
	buffer := make([]byte,512)
	_,err := file.Read(buffer)
	if err != nil {
		return "",err
	}
	ContentType := http.DetectContentType(buffer)

	return ContentType,nil
}