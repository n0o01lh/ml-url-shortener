package data

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/ml-url-shortener/internals/core/domain"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
)

type DynamoDb struct {
	client *dynamodb.Client
}

func NewDynamoDb(client *dynamodb.Client) *DynamoDb {
	return &DynamoDb{
		client: client,
	}
}

var _ ports.DB = (*DynamoDb)(nil)

func (d *DynamoDb) List() {
	resp, err := d.client.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	for _, tableName := range resp.TableNames {
		fmt.Println(tableName)
	}
}

func (d *DynamoDb) PutUrl(shortedUrl *domain.ShortedUrl) error {
	payload := map[string]types.AttributeValue{
		"id":           &types.AttributeValueMemberS{Value: shortedUrl.Id},
		"original_url": &types.AttributeValueMemberS{Value: shortedUrl.OriginalUrl},
		"available":    &types.AttributeValueMemberBOOL{Value: *shortedUrl.Available},
		"created_at":   &types.AttributeValueMemberS{Value: shortedUrl.CreatedAt},
		"updated_at":   &types.AttributeValueMemberS{Value: shortedUrl.UpdatedAt},
	}

	_, err := d.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("url"),
		Item:      payload,
	})

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (d *DynamoDb) UpdateUrl(key string, shortedUrl *domain.ShortedUrl) error {

	payload := make(map[string]types.AttributeValue)
	payload[":updatedAt"] = &types.AttributeValueMemberS{Value: shortedUrl.UpdatedAt}
	updateExpression := "SET updated_at = :updatedAt"

	if shortedUrl.OriginalUrl != "" {
		payload[":originalUrl"] = &types.AttributeValueMemberS{Value: shortedUrl.OriginalUrl}
		updateExpression = updateExpression + ", original_url = :originalUrl"
	}

	if shortedUrl.Available != nil {
		payload[":available"] = &types.AttributeValueMemberBOOL{Value: *shortedUrl.Available}
		updateExpression = updateExpression + ", available = :available"
	}

	_, err := d.client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String("url"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: key},
		},
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: payload,
		ReturnValues:              types.ReturnValueUpdatedNew,
	})

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (d *DynamoDb) GetUrl(key string) (*domain.ShortedUrl, error) {

	result, err := d.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("url"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: key},
		},
	})

	if err != nil {
		log.Error(err)
		return nil, err
	}

	if result.Item == nil {
		fmt.Println("Item not found")
		return nil, err
	}

	var shortedUrl *domain.ShortedUrl
	err = attributevalue.UnmarshalMap(result.Item, &shortedUrl)

	if original_url, ok := result.Item["original_url"].(*types.AttributeValueMemberS); ok {
		shortedUrl.OriginalUrl = original_url.Value
	}

	if err != nil {
		log.Fatalf("Failed to unmarshal record: %v", err)
	}

	return shortedUrl, nil
}

func (d *DynamoDb) PutStats(stats *domain.Stats) error {

	payload := map[string]types.AttributeValue{
		"id":     &types.AttributeValueMemberS{Value: stats.Id},
		"clicks": &types.AttributeValueMemberN{Value: *aws.String(strconv.FormatInt(stats.Clicks, 10))},
	}

	_, err := d.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("stats"),
		Item:      payload,
	})

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (d *DynamoDb) UpdateStats(key string) error {

	payload := map[string]types.AttributeValue{
		":click": &types.AttributeValueMemberN{Value: "1"},
	}

	_, err := d.client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String("stats"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: key},
		},
		UpdateExpression:          aws.String("SET clicks = clicks + :click"),
		ExpressionAttributeValues: payload,
		ReturnValues:              types.ReturnValueUpdatedNew,
	})

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (d *DynamoDb) GetStats(key string) (*domain.Stats, error) {
	result, err := d.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("stats"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: key},
		},
	})

	if err != nil {
		log.Error(err)
		return nil, err
	}

	if result.Item == nil {
		fmt.Println("Item not found")
		return nil, err
	}

	var stats *domain.Stats
	err = attributevalue.UnmarshalMap(result.Item, &stats)

	if err != nil {
		log.Fatalf("Failed to unmarshal record: %v", err)
	}

	return stats, nil
}
