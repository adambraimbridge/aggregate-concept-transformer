package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/coreos/fleet/log"
	"io"
	"strings"
	"net/http"
	"net"
	"time"
)

type Client struct {
	s3         *s3.S3
	bucketName string
}

func NewClient(bucketName string, awsRegion string) (Client, error) {
	hc := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          20,
			IdleConnTimeout:       90 * time.Second,
			MaxIdleConnsPerHost:   20,
			TLSHandshakeTimeout:   3 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	sess, err := session.NewSession(
		&aws.Config{
			Region:     aws.String(awsRegion),
			MaxRetries: aws.Int(1),
			HTTPClient: &hc,
		})
	if err != nil {
		return Client{}, err
	}
	client := s3.New(sess)

	return Client{
		s3:         client,
		bucketName: bucketName,
	}, err
}

func (c *Client) GetConcept(UUID string) (bool, io.ReadCloser, error) {
	params := &s3.GetObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(getKey(UUID)),
	}

	resp, err := c.s3.GetObject(params)
	if err != nil {
		e, ok := err.(awserr.Error)
		if ok && e.Code() == "NoSuchKey" {
			return false, nil, nil
		}
		return false, nil, err
	}
	return true, resp.Body, err
}

func (c *Client) HealthCheck() (string, error) {
	params := &s3.HeadBucketInput{
		Bucket: aws.String(c.bucketName), // Required
	}
	_, err := c.s3.HeadBucket(params)
	if err != nil {
		log.Errorf("Got error running S3 health check, %v", err.Error())
		return "Can not perform check on S3 bucket", err
	}
	return "Access to S3 bucket ok", err
}

func getKey(UUID string) string {
	return strings.Replace(UUID, "-", "/", -1)
}
