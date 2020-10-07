/*
 * File Name FireStorageImpl.go
 * Created on Sun Oct 04 2020
 *
 * Copyright (c) 2020
 */

package infrastructure

import (
	"context"
	"io"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)
 
type FireStorageImpl struct {

}

func NewFireStorageImpl() IFireStorage {
	return &FireStorageImpl{}
}

func (f *FireStorageImpl) NewApp() (*firebase.App,error) {
	CredentialFiles :=  option.WithCredentialsFile("keystore.json")
	FirebaseNewApp,err := firebase.NewApp(context.Background(),&firebase.Config{
		StorageBucket: "backend-jobs-go.appspot.com",
	},CredentialFiles)
	
	if err != nil {
		return nil,err
	}
	
	return FirebaseNewApp,err
}

func (f *FireStorageImpl) UploadPhotoProfile(ObjectName,FilePath string) (string,error) {
	
	client,err := f.NewApp()
	if err != nil {
		return "",err
	}
	
	storage,err := client.Storage(context.Background())
	if err != nil {
		return "",err
	}

	bucket,err := storage.DefaultBucket() 
	if err != nil {
		return "",err
	}
	
	file,err := os.Open(FilePath)
	if err != nil {
		return "",err
	}
	defer file.Close()
	
	wc := bucket.Object("employe-photo-"+ObjectName).NewWriter(context.Background())
	if _,err := io.Copy(wc,file); err != nil {
		return "",err
	}
	
	if err := wc.Close(); err != nil {
		return "",err
	}

	// Berfungsi untuk mendapatkan link file atau Download
	GetUrl,err := bucket.Object("employe-photo-"+ObjectName).Attrs(context.Background())
	if err != nil {
		return "",err
	}

	return GetUrl.MediaLink,nil
}
