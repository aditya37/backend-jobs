/*
 * File Name IFireStorage.go
 * Created on Sun Oct 04 2020
 *
 * Copyright (c) 2020
 */

package infrastructure

type IFireStorage interface {
	UploadPhotoProfile(ObjectName,FilePath string) (string,error)
}