package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common/asyncq"
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common/utility"
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/tasks"
)

type UploadImageForm struct {
	PageTitle  string
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

func DisplayUploadImageForm(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {
	RenderGatedTemplate(w, WebAppRoot+"/templates/uploadimageform.html", u)
}

func ProcessUploadImage(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {

	shouldProcessThumbnailAsynchronously := true

	file, fileheader, err := r.FormFile("imagefile")

	if err != nil {
		log.Println("Encountered error when attempting to read uploaded file: ", err)
	}

	randomFileName := utility.GenerateUUID()

	if fileheader != nil {

		extension := filepath.Ext(fileheader.Filename)
		_ = r.ParseMultipartForm(32 << 20)

		defer file.Close()

		imageFilePathWithoutExtension := "./static/uploads/images/" + randomFileName
		f, err := os.OpenFile(imageFilePathWithoutExtension+extension, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			log.Println(err)
			return
		}

		defer f.Close()
		_, _ = io.Copy(f, file)

		// Note: Moved the thumbnail generation logic (commented out code block below) to the
		// ImageResizeTask object in the tasks package.
		thumbnailResizeTask := tasks.NewImageResizeTask(imageFilePathWithoutExtension, extension)

		if shouldProcessThumbnailAsynchronously {

			asyncq.TaskQueue <- thumbnailResizeTask

		} else {

			thumbnailResizeTask.Perform()
		}

		m := make(map[string]string)
		m["thumbnailPath"] = strings.TrimPrefix(imageFilePathWithoutExtension, ".") + "_thumb.png"
		m["imagePath"] = strings.TrimPrefix(imageFilePathWithoutExtension, ".") + ".png"
		m["PageTitle"] = "Image Preview"

		RenderGatedTemplate(w, WebAppRoot+"/templates/imagepreview.html", m)

	} else {
		_, _ = w.Write([]byte("Failed to process uploaded file!"))
	}
}

func ValidateUploadImageForm(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {

	ProcessUploadImage(w, r, u)

}

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {

	u := UploadImageForm{}
	u.Fields = make(map[string]string)
	u.Errors = make(map[string]string)
	u.PageTitle = "Upload Image"

	switch r.Method {

	case "GET":
		DisplayUploadImageForm(w, r, &u)
	case "POST":
		ValidateUploadImageForm(w, r, &u)
	default:
		DisplayUploadImageForm(w, r, &u)
	}

}
