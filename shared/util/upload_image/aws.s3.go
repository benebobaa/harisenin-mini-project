package upload_image

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	awsSess "github.com/aws/aws-sdk-go/aws/session"
	"github.com/benebobaa/harisenin-mini-project/shared/util"
)

func NewSessionAWSS3(c util.Config) (*awsSess.Session, error) {

	sess, err := awsSess.NewSession(&aws.Config{
		Region:      aws.String(c.AWSRegion),
		Credentials: credentials.NewStaticCredentials(c.AWSAccessKey, c.AWSSecretKey, ""),
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}
