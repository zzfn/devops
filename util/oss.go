package util

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
)

var endpoint = "endpoint"
var accessKeyID = "accessKeyID"
var accessKeySecret = "accessKeySecret"

func PutOss() {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		print(3, err)
		// HandleError(err)
	}

	bucket, err := client.Bucket("zzfzzf")
	if err != nil {
		print(2, err)
		// HandleError(err)
	}
	getFile(BasePath+"/workspace/react-webpack/dist", "test", bucket)
}
func getFile(workSpacePath, ossPrefix string, bucket *oss.Bucket) {
	files, _ := ioutil.ReadDir(workSpacePath)
	for _, file := range files {
		if file.IsDir() {
			getFile(workSpacePath+"/"+file.Name(), ossPrefix+"/"+file.Name(), bucket)
		} else {
			err := bucket.PutObjectFromFile(ossPrefix+"/"+file.Name(), workSpacePath+"/"+file.Name())
			if err != nil {
				return
			}
		}
	}
}
