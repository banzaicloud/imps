package ecr

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type StringableCredentials struct {
	aws.Credentials
	// Region specifies which region to connect to when using this credential
	Region string
}

func (c StringableCredentials) String() string {
	return fmt.Sprintf("%s/%s/%s/%s", c.Region, c.AccessKeyID, c.SecretAccessKey, c.SessionToken)
}
