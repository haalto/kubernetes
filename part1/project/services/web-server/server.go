package main

import (
	"errors"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func dowloadNewImageAndSaveItToVolumeAndStatics() error {
	response, err := http.Get("https://picsum.photos/1200")
	if err != nil {
		log.Print(err)
		return err
	}

	img, err := jpeg.Decode(response.Body)

	if (err != nil) {
		log.Print(err)
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Print(err)
		return errors.New("received non 200 response code")
	}

	fileVolume, err := os.Create(filepath.Join("/", "usr", "src", "app", "files", "img.jpg"))
	if err != nil {
		log.Print(err)
		return err
	}
	defer fileVolume.Close()

	err = jpeg.Encode(fileVolume, img, nil)
	if err != nil {
		log.Print(err)
		return err
	}

	copyFileToStatic()
	return nil
}

func copyFileToStatic() error {
	in, err := os.Open(filepath.Join("/", "usr", "src", "app", "files", "img.jpg"))
	if err != nil {
			return err
	}
	defer in.Close()

	out, err := os.Create(filepath.Join("static", "img.jpg"))
	if err != nil {
			return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
			return err
	}
	return out.Close()
}

func checkIfFileExistsOnVolume() bool {
	_, err := os.Stat(filepath.Join("/", "usr", "src", "app", "files", "img.jpg"))
	if err == nil {
			return true
	}
	if errors.Is(err, os.ErrNotExist) {
			return false
	}
	return false
}

func updateImgAtTime() {
	for {
		t := time.Now()
		if t.Hour() == 0 && t.Minute() == 0 && t.Second() == 0 {
			dowloadNewImageAndSaveItToVolumeAndStatics()
		}
		time.Sleep(time.Second)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	if checkIfFileExistsOnVolume() {
		copyFileToStatic()
	} else {
		dowloadNewImageAndSaveItToVolumeAndStatics()
	}

	go updateImgAtTime()

	fmt.Println("Server started in 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}