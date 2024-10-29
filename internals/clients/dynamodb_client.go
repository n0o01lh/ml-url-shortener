package clients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoDbClient(awsAccessKey, awsSecretAccessKey string) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-2"),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: awsAccessKey, SecretAccessKey: awsSecretAccessKey, SessionToken: "",
			},
		}),
	)
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)
	return client, nil
}
