package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common/utility"
)

type UploadVideoForm struct {
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

func DisplayUploadVideoForm(w http.ResponseWriter, r *http.Request, u *UploadVideoForm) {
	RenderTemplate(w, WebAppRoot+"/templates/uploadvideoform.html", u)
}

func ProcessUploadVideo(w http.ResponseWriter, r *http.Request, u *UploadVideoForm) {

	file, fileheader, err := r.FormFile("videofile")

	if err != nil {
		log.Println("Encountered error when attempting to read uploaded file: ", err)
	}

	randomFileName := utility.GenerateUUID()

	if fileheader != nil {

		extension := filepath.Ext(fileheader.Filename)
		_ = r.ParseMultipartForm(32 << 20)

		defer file.Close()

		videoFilePathWithoutExtension := "./static/uploads/videos/" + randomFileName
		f, err := os.OpenFile(videoFilePathWithoutExtension+extension, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			log.Println(err)
			return
		}

		defer f.Close()
		_, _ = io.Copy(f, file)

		thumbImageFilePath := videoFilePathWithoutExtension + "_thumb.png"
		command := "ffmpeg -y  -i " + videoFilePathWithoutExtension + extension + " -s 430x372 -f mjpeg -vframes 1 -ss 3 " + thumbImageFilePath
		commandOutput, err := exec.Command("sh", "-c", command).Output()
		fmt.Println("video thumbnail generation command output: ", commandOutput)

		if err != nil {
			log.Println(err)
			return
		}

		m := make(map[string]string)
		m["thumbnailPath"] = strings.TrimPrefix(videoFilePathWithoutExtension, ".") + "_thumb.png"
		m["videoPath"] = strings.TrimPrefix(videoFilePathWithoutExtension, ".") + ".mp4"

		RenderTemplate(w, WebAppRoot+"/templates/videopreview.html", m)

	} else {
		_, _ = w.Write([]byte("Failed to process uploaded file!"))
	}
}

func ValidateUploadVideoForm(w http.ResponseWriter, r *http.Request, u *UploadVideoForm) {

	ProcessUploadVideo(w, r, u)

}

func UploadVideoHandler(w http.ResponseWriter, r *http.Request) {

	u := UploadVideoForm{}
	u.Fields = make(map[string]string)
	u.Errors = make(map[string]string)

	switch r.Method {

	case "GET":
		DisplayUploadVideoForm(w, r, &u)
	case "POST":
		ValidateUploadVideoForm(w, r, &u)
	default:
		DisplayUploadVideoForm(w, r, &u)
	}

}
