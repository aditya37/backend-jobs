/*
 * File Name IFireStorage.go
 * Created on Sun Oct 04 2020
 *
 * Copyright (c) 2020
 */

package infrastructure

type IFireStorage interface {
	UploadPhotoProfile(ObjectName,FilePath string) (string,error)
	UploadPortofolio(ObjectName,FilePath string) (string,string,error)
	UploadResume(ObjectName,FilePath string) (string,string,error)
	GetUrl(ObjectName string) (string,error)
}