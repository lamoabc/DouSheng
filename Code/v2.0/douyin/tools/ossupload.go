package tools

import (
	"fmt"
	"mime/multipart"
	"os"
)

func GetPlayUrl(filename string, content multipart.File) string {
	bucket := OssUpload()
	playurl := "Video address/" + filename
	// 将视频传到对应的文件夹
	err := bucket.PutObject(playurl, content)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return "https://yygh-lamo.oss-cn-beijing.aliyuncs.com/" + playurl
}

func GetCoverUrl(filename string) string {
	//bucket := OssUpload()
	//objectName := "Video address/" + filename
	//targetBucketName := "yygh-lamo"
	//style := "?x-oss-process=video/snapshot,t_0,f_jpg,w_0,h_0,m_fast,ar_auto"

	return "https://yygh-lamo.oss-cn-beijing.aliyuncs.com/"
}
