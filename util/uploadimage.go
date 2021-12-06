package util

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"member-service/exception"
	"member-service/model"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func ObjectStorageService(bucketprofile *oss.Bucket, id string, profilepicture string, domain string, dir string, act string) (response model.UploadResponse) {

	id = strings.ReplaceAll(id, " ", "")
	nowInt := time.Now().UnixNano() / int64(time.Millisecond)

	/**Create PNG temporary file*/
	png := toPng(id, profilepicture)
	/**Is Created*/
	ext := ".png"

	if png == "png: invalid format: not a PNG file" {
		/**Create JPEG temporary file*/
		toJpeg(id, profilepicture)
		ext = ".jpeg"
		/**Is Created*/
	}

	object := ""
	if act == "TEST" {
		object = "profile/test/" + strconv.FormatInt(nowInt, 10) + "-" + id + ext
	} else {
		object = "profile/" + strconv.FormatInt(nowInt, 10) + "-" + id + ext
	}

	localFile := dir + id + ext

	err := bucketprofile.PutObjectFromFile(object, localFile)
	if err != nil {
		err = os.Remove(localFile) //Remove temporary file
		exception.PanicIfNeeded(err)
	}
	/**Completed Upload*/

	err = os.Remove(localFile) //Remove temporary file
	exception.PanicIfNeeded(err)

	response = model.UploadResponse{
		Object: object,
		Domain: domain,
	}
	return response
}

/**Create PNG temporary file*/
func toPng(id string, profilepicture string) string {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(strings.Split(profilepicture, "base64,")[1]))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		return err.Error()
	}
	bounds := m.Bounds()
	fmt.Println(bounds, formatString)

	//Encode from image format to writer
	pngFilename := os.Getenv("FILE_DIR") + id + ".png"
	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err.Error()
	}

	err = png.Encode(f, m)
	if err != nil {
		return err.Error()
	}
	return "true"
}

/**Create JPEG temporary file*/
func toJpeg(id string, profilepicture string) {
	unbased, err := base64.StdEncoding.DecodeString((strings.Split(profilepicture, "base64,")[1]))
	exception.PanicIfNeeded(err)

	r := bytes.NewReader(unbased)
	im, err := jpeg.Decode(r)
	exception.PanicIfNeeded(err)

	f, err := os.OpenFile(os.Getenv("FILE_DIR")+id+".jpeg", os.O_WRONLY|os.O_CREATE, 0777)
	exception.PanicIfNeeded(err)

	jpeg.Encode(f, im, &jpeg.Options{
		Quality: 100,
	})
}
