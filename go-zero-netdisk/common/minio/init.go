package minio

import (
	"github.com/minio/minio-go"
	"log"
	"path"
	"strconv"
	"strings"
	"time"
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
		BucketName string
		client     *minio.Client
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

	return &Client{bucketName, client}
}

// GenObjectName ext can only be one string variable
func (c *Client) GenObjectName(hash, rawFilename string, ext ...string) (string, string) {
	var filename, e string
	if len(ext) > 0 {
		e = ext[0]
		filename = strings.TrimSuffix(rawFilename, e)
	} else {
		e = path.Ext(rawFilename)
		filename = strings.TrimSuffix(rawFilename, e)
	}

	filename += "|" + strconv.FormatInt(time.Now().Unix(), 10) + e
	return filename, "/" + time.Now().Format("2006-01") + "/" +
		string(hash[0]) + "/" + string(hash[0]) + "/" + filename
}
