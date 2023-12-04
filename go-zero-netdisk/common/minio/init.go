package minio

import (
	"github.com/minio/minio-go"
	"log"
)

type (
	Conf struct {
		Endpoint        string
		AccessKeyId     string
		SecretAccessKey string
		BucketName      string
		UseSSL          bool
	}

	Client struct {
		client *minio.Client
	}
)

// Init 初使化 minio client对象。
func Init(conf *Conf) *Client {
	client, err := minio.New(conf.Endpoint, conf.AccessKeyId, conf.SecretAccessKey, conf.UseSSL)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	bucketName := conf.BucketName
	exists, err := client.BucketExists(bucketName)
	if err != nil || !exists {
		err = client.MakeBucket(bucketName, "")
		if err != nil {
			log.Fatalln(err)
		}
	}

	return &Client{client}
}
