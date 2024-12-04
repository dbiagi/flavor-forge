package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dbiagi/gororoba/src/domain"
	"log/slog"
)

func CreateDynamoDBConnection(awsConfig AWSConfig) (*dynamodb.DynamoDB, *domain.Error) {
	config := aws.Config{
		Region: aws.String(awsConfig.Region),
	}
	sess := session.Must(session.NewSession(&config))
	dynamoDB := dynamodb.New(sess)
	_, err := dynamoDB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		slog.Error("Error creating session on AWS", slog.String("error", err.Error()))
		return nil, &domain.Error{
			Message: "Error creating session on AWS. Message: " + err.Error(),
			Cause:   err,
		}
	}

	return dynamoDB, nil
}
