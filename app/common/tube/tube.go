package tube

import (
	"AILN/app/common"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"

	"strconv"
	"time"
)

func getToken() string {
	var maxInt uint64 = 1 << 32
	putPolicy := storage.PutPolicy{
		Scope:   common.CONFIG.String("oss.bucket"),
		Expires: maxInt,
	}
	mac := qbox.NewMac(common.CONFIG.String("oss.accessKey"), common.CONFIG.String("oss.secretKey"))
	return putPolicy.UploadToken(mac)
}

func getObjectName(filename string, id uint) (string, error) {
	i := strings.LastIndex(filename, ".")
	fileType := filename[i+1:]

	timeEpochNow := time.Now().Unix()
	objectName := strconv.FormatUint(uint64(id), 10) + "-" + strconv.FormatInt(timeEpochNow, 10) + "." + fileType
	return objectName, nil
}

func UploadFile(filename string, id uint, r io.ReaderAt, dataLen int64) (string, error) {
	upToken := getToken()
	objectName, err := getObjectName(filename, id)
	if err != nil {
		return "", err
	}

	// 下面是七牛云的oss所需信息，objectName对应key是文件上传路径
	cfg := storage.Config{Zone: &storage.ZoneHuanan, UseHTTPS: false, UseCdnDomains: true}
	formUploader := storage.NewResumeUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputExtra{Params: map[string]string{"x:name": "lablab"}}
	err = formUploader.Put(context.Background(), &ret, upToken, objectName, r, dataLen, &putExtra)
	if err != nil {
		return "", err
	}

	url := "https://" + common.CONFIG.String("oss.domain") + "/" + objectName
	return url, nil
}

func Download(url string) string {
	index := strings.LastIndex(url, "/")
	key := url[index+1:]
	domain := url[:index]
	fmt.Println(key, domain)
	mac := qbox.NewMac(common.CONFIG.String("oss.accessKey"), common.CONFIG.String("oss.secretKey"))
	deadline := time.Now().Add(time.Second * 3600).Unix() // 1小时有效期

	return storage.MakePrivateURL(mac, domain, key, deadline)
}
