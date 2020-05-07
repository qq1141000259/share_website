package utils

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	cos "github.com/tencentyun/cos-go-sdk-v5"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
)

type CosClient struct {
	*cos.Client
	Bk    string
	ScId  string
	ScKey string
}

func NewCos(bucket, scId, scKey string) *CosClient {
	u, _ := url.Parse(bucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  scId,
			SecretKey: scKey,
		},
	})
	return &CosClient{client, bucket, scId, scKey}
}

// 上传byte内容
func (c *CosClient) PutBytes(fileName string, file io.Reader) error {
	_, err := c.Object.Put(context.Background(), fileName, file, nil)
	if err != nil {
		return err
	}
	return nil
}

// 本地上传
func (c *CosClient) LocalUpload(fileName, filePath string) error {
	_, err := c.Object.PutFromFile(context.Background(), fileName, "../test", nil)
	if err != nil {
		return err
	}
	return nil
}

// 获取下载链接
func (c *CosClient) GetDownloadUrl(fileName string) (string, error) {
	presignedURL, err := c.Object.GetPresignedURL(context.Background(), http.MethodGet, fileName, c.ScId, c.ScKey, time.Hour, nil)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

// sts sdk
func (c *CosClient) CreateCredentials() (*sts.Credentials, error) {
	s := sts.NewClient(
		c.ScId,
		c.ScKey,
		nil,
	)
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          "ap-shanghai",
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
						"name/cos:PutObject",
					},
					Effect: "allow",
					Resource: []string{
						"*",
					},
				},
			},
		},
	}
	res, err := s.GetCredential(opt)
	if err != nil {
		return nil, err
	}
	return res.Credentials, nil
}
