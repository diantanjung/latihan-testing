package thirdpartylib

import (
	"context"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3GetObjectAPI interface {
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

type UploaderS3 struct {
	api S3GetObjectAPI
}

func (u UploaderS3) GetObjectFromS3(ctx context.Context, bucket, key string) ([]byte, error) {
	object, err := u.api.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	defer object.Body.Close()

	return ioutil.ReadAll(object.Body)
}
