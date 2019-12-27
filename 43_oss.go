package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {
	// 创建OSSClient实例。
	client, err := oss.New("oss-cn-hangzhou", "LTAI4FirCFtGKJWqFJNDJFJe", "VydHHa5wHWa01KNnV0fWboAgieM5F7")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket("learnta-backup")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 指定存储类型为标准存储，缺省也为标准存储。
	storageType := oss.ObjectStorageClass(oss.StorageStandard)

	// 指定存储类型为归档存储。
	// storageType := oss.ObjectStorageClass(oss.StorageArchive)

	// 指定访问权限为公共读，缺省为继承bucket的权限。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	// 上传字符串。
	err = bucket.PutObject("test01.json", strings.NewReader("yourObjectValue"), storageType, objectAcl)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
