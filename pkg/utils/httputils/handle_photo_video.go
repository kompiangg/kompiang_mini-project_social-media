package httputils

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"time"

	"github.com/kompiang_mini-project_social-media/config"
	"github.com/kompiang_mini-project_social-media/pkg/utils/fileutils"
)

// func HandlePhotoAndVideoForm(image *multipart.FileHeader, video *multipart.FileHeader) (*string, *string, error) {
// 	fileutils.CheckUploadFolder()
// 	imageFilename := path.Join(config.GetConfig().UploadFolderName, fmt.Sprintf("%d-%s", time.Now().UnixMilli(), image.Filename))
// 	videoFilename := path.Join(config.GetConfig().UploadFolderName, fmt.Sprintf("%d-%s", time.Now().UnixMilli(), video.Filename))

// 	errChan := make(chan error)
// 	defer close(errChan)

// 	// Image
// 	go func(errChan chan error) {
// 		imageSrc, err := image.Open()
// 		if err != nil {
// 			log.Println("[HandlePhotoAndVideoForm]", err.Error())
// 			errChan <- err
// 			return
// 		}
// 		defer imageSrc.Close()

// 		imageDst, err := os.Create(imageFilename)
// 		if err != nil {
// 			log.Println("[HandlePhotoAndVideoForm]", err.Error())
// 			errChan <- err
// 			return
// 		}

// 		_, err = io.Copy(imageDst, imageSrc)
// 		if err != nil {
// 			log.Println("[HandlePhotoAndVideoForm]", err.Error())
// 			errChan <- err
// 			return
// 		}

// 		errChan <- nil
// 	}(errChan)

// 	// Video
// 	go func(errChan chan error) {
// 		videoSrc, err := video.Open()
// 		if err != nil {
// 			log.Println("[HandlePhotoAndVideoForm]", err.Error())
// 			errChan <- err
// 			return
// 		}
// 		defer videoSrc.Close()

// 		videoDst, err := os.Create(videoFilename)
// 		if err != nil {
// 			log.Println("[HandlePhotoAndVideoForm]", err.Error())
// 			errChan <- err
// 			return
// 		}

// 		_, err = io.Copy(videoDst, videoSrc)
// 		if err != nil {
// 			log.Println("[HandlePhotoAndVideoForm]", err.Error())
// 			errChan <- err
// 			return
// 		}

// 		errChan <- nil
// 	}(errChan)

// 	for i := 0; i < 2; i++ {
// 		err := <-errChan
// 		if err != nil {
// 			os.Remove(imageFilename)
// 			os.Remove(videoFilename)
// 			return nil, nil, err
// 		}
// 	}

// 	return &imageFilename, &videoFilename, nil
// }

func HandleFileForm(file *multipart.FileHeader) (*string, error) {
	fileutils.CheckUploadFolder()
	filename := path.Join(config.GetConfig().UploadFolderName, fmt.Sprintf("%d-%s", time.Now().UnixMilli(), file.Filename))

	fileSrc, err := file.Open()
	if err != nil {
		log.Println("[HandleFileForm]", err.Error())
		return nil, err
	}
	defer fileSrc.Close()

	fileDst, err := os.Create(filename)
	if err != nil {
		log.Println("[HandleFileForm]", err.Error())
		os.Remove(filename)
		return nil, err
	}
	defer fileDst.Close()

	_, err = io.Copy(fileDst, fileSrc)
	if err != nil {
		log.Println("[HandleFileForm]", err.Error())
		os.Remove(filename)
		return nil, err
	}

	return &filename, nil
}

// func HandleVideoForm(video *multipart.FileHeader) (*string, error) {
// 	fileutils.CheckUploadFolder()
// 	videoFilename := path.Join(config.GetConfig().UploadFolderName, fmt.Sprintf("%d-%s", time.Now().UnixMilli(), video.Filename))

// 	videoSrc, err := video.Open()
// 	if err != nil {
// 		log.Println("[HandlePhotoAndVideoForm]", err.Error())
// 		return nil, err
// 	}
// 	defer videoSrc.Close()

// 	videoDst, err := os.Create(videoFilename)
// 	if err != nil {
// 		log.Println("[HandlePhotoAndVideoForm]", err.Error())
// 		os.Remove(videoFilename)
// 		return nil, err
// 	}
// 	defer videoDst.Close()

// 	_, err = io.Copy(videoDst, videoSrc)
// 	if err != nil {
// 		log.Println("[HandlePhotoAndVideoForm]", err.Error())
// 		os.Remove(videoFilename)
// 		return nil, err
// 	}

// 	return &videoFilename, nil
// }
