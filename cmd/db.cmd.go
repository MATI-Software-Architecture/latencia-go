package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mati/latencia/schema"
)

// https://medium.com/yemeksepeti-teknoloji/dynamodb-with-aws-sdk-go-v2-part-2-crud-operations-3da68c2f431f

const (
	TableName      = "reto1"
	Location       = "us-east-2"
	Api_key_id     = "API_KEY"
	Secret_api_key = "SECRET_API_KEY"
	Session_token  = ""
)

func InsertItem(item schema.TableStruct) error {
	data, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				Api_key_id,
				Secret_api_key,
				Session_token,
			),
		),
		config.WithRegion(Location),
	)
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item:      data,
	})
	if err != nil {
		return err
	}
	return nil
}
