package util

import (
	"member-service/exception"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func DeleteImage(bucketprofile *oss.Bucket, id string, object string) string {

	id = strings.ReplaceAll(id, " ", "")

	err := bucketprofile.DeleteObject(object)
	if err != nil {
		exception.PanicIfNeeded(err)
	}
	/**Completed Upload*/

	return "Berhasil hapus foto"
}
