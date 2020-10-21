/*
 * File Name FireStorageImpl.go
 * Created on Sun Oct 04 2020
 *
 * Copyright (c) 2020
 */

package infrastructure

import (
	"context"
	"fmt"
	"io"
	"os"

	firebase "firebase.google.com/go"
	"github.com/google/uuid"
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
	
	id := uuid.New()
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
	wc.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens":id.String()}
	if _,err := io.Copy(wc,file); err != nil {
		return "",err
	}
	
	if err := wc.Close(); err != nil {
		return "",err
	}
	URL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/backend-jobs-go.appspot.com/o/%s?alt=media&token=%s",wc.Name,id.String())
	return URL,nil
}
// FIXME: Buat Menjadi pisah, tidak bisa di buat multiple ex: Portofolio upload
func (f *FireStorageImpl) UploadPortofolio(ObjectName,FilePath string) (string,string,error) {
	
	id := uuid.New()

	Client,err := f.NewApp()
	if err != nil {
		return "","",err
	}

	Storage,err := Client.Storage(context.Background())
	if err != nil {
		return "","",err
	}
	Bucket,err := Storage.DefaultBucket()
	if err != nil {
		return "","",err
	}

	file,err := os.Open(FilePath)
	if err != nil {
		return "","",err
	}
	defer file.Close()

	wc := Bucket.Object("employe-portofolio-"+ObjectName).NewWriter(context.Background())
	wc.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens":id.String()}
	if _,err := io.Copy(wc,file); err != nil {
		return "","",err
	}
	if err := wc.Close();err != nil {
		return "","",err
	}
	URL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/backend-jobs-go.appspot.com/o/%s?alt=media&token=%s",wc.Name,id.String())
	return URL,wc.Name,nil
}

func (f *FireStorageImpl) GetUrl(ObjectName string) (string,error) {

	Client,err := f.NewApp()
	if err != nil {
		return "",err
	}
	Storage,err := Client.Storage(context.Background())
	if err != nil {
		return "",err
	}
	Bucket,err := Storage.DefaultBucket()
	if err != nil {
		return "",err
	}
	GetUrl,err := Bucket.Object(ObjectName).Attrs(context.Background())
	if err != nil {
		return "",err
	}

	return GetUrl.MediaLink,nil
}
func (f *FireStorageImpl) UploadResume(ObjectName,FilePath string) (string,string,error) {
	
	id := uuid.New() 

	Client,err := f.NewApp()
	if err != nil {
		return "","",err
	}
	Storage,err := Client.Storage(context.Background())
	if err != nil {
		return "","",err
	}
	Bucket,err := Storage.DefaultBucket()
	if err != nil {
		return "","",err
	}

	file,err := os.Open(FilePath)
	if err != nil {
		return "","",err
	}
	defer file.Close()
	
	wc := Bucket.Object("employe-resume-"+ObjectName).NewWriter(context.Background())
	wc.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens":id.String()}
	if _,err := io.Copy(wc,file); err != nil {
		return "","",err
	}
	if err := wc.Close(); err != nil {
		return "","",err
	}

	URL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/backend-jobs-go.appspot.com/o/%s?alt=media&token=%s",wc.Name,id.String())

	return URL,wc.Name,nil
}