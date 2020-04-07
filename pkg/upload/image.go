package upload

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"um_sys/pkg/file"
	"um_sys/pkg/setting"
	"um_sys/pkg/util"
)

func GetImageFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/api/chat/" + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := file.GetExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(ext) == strings.ToUpper(allowExt) {
			return true
		}
	}
	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

func PrepareCreateImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return err
	}

	noPerm := file.CheckNoPermission(src)
	if noPerm {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
