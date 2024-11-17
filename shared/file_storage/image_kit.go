package filestorage

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"sync"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
)

var instance *imagekit.ImageKit

func Init() {
	instance = imagekit.NewFromParams(imagekit.NewParams{
		PrivateKey: envs.privateKey,
		PublicKey: envs.publicKey,
		UrlEndpoint: envs.endpointURL,
	})
	
	if instance == nil {
		log.Fatal("File Storage failed to connect!")
	}

	fmt.Println("File storage connected succesfully!")
}

func UploadFile(file multipart.File, folder string) (data *uploader.UploadResult, err error) {
	response, err := instance.Uploader.Upload(
		context.Background(),
		file,
		uploader.UploadParam{
			Folder: folder,
			FileName: "image",
		},
	)
	if err != nil {
		return nil, err
	}
	return &response.Data, err
}

func UploadFiles(files []multipart.File, folder string) (upload []*uploader.UploadResult, errs []string) {
	results := make([]*uploader.UploadResult, len(files))
	errors := make([]string, len(files))

	var wg sync.WaitGroup
	wg.Add(len(files))
	
	for i, file := range files {
		go func() {
			defer wg.Done()
			response, err := instance.Uploader.Upload(
				context.Background(),
				file,
				uploader.UploadParam{
					Folder:   folder,
					FileName: "image",
				},
			)
			if err != nil {
				errors[i] = err.Error()
				return
			}
			results[i] = &response.Data
		}()
	}

	wg.Wait()
	return results, errs
}

func DeleteFile(id string) error {
	_, err := instance.Media.DeleteFile(context.Background(), id)
	return err
}