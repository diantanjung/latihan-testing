package thirdpartylib

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUploader struct {
	mock.Mock
}

func (m MockUploader) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*s3.GetObjectOutput), args.Error(1)
}

func TestGetObjectFromS3_Mock(t *testing.T) {
	m := new(MockUploader)
	u := UploaderS3{
		api: m,
	}

	m.On("GetObject", nil, nil).Return(nil, true)

	output := &s3.GetObjectOutput{
		Body: ioutil.NopCloser(bytes.NewReader([]byte("this is the body foo bar baz"))),
	}
	m.On("GetObject", mock.Anything, mock.Anything).Return(output, nil)

	ctx := context.TODO()
	content, err := u.GetObjectFromS3(ctx, "fooBucket", "barKey")

	assert.NoError(t, err)

	expect := []byte("this is the body foo bar baz")
	assert.Equal(t, content, expect)
}
