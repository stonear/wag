package dynamodb

import (
	"context"
	"fmt"
	"strings"

	"github.com/Clever/wag/samples/gen-go-db/models"
	"github.com/Clever/wag/samples/gen-go-db/server/db"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/go-openapi/strfmt"
)

var _ = strfmt.DateTime{}

// NoRangeThingWithCompositeAttributesTable represents the user-configurable properties of the NoRangeThingWithCompositeAttributes table.
type NoRangeThingWithCompositeAttributesTable struct {
	DynamoDBAPI        dynamodbiface.DynamoDBAPI
	Prefix             string
	TableName          string
	ReadCapacityUnits  int64
	WriteCapacityUnits int64
}

// ddbNoRangeThingWithCompositeAttributesPrimaryKey represents the primary key of a NoRangeThingWithCompositeAttributes in DynamoDB.
type ddbNoRangeThingWithCompositeAttributesPrimaryKey struct {
	NameBranch string `dynamodbav:"name_branch"`
}

// ddbNoRangeThingWithCompositeAttributesGSINameVersion represents the nameVersion GSI.
type ddbNoRangeThingWithCompositeAttributesGSINameVersion struct {
	NameVersion string          `dynamodbav:"name_version"`
	Date        strfmt.DateTime `dynamodbav:"date"`
}

// ddbNoRangeThingWithCompositeAttributes represents a NoRangeThingWithCompositeAttributes as stored in DynamoDB.
type ddbNoRangeThingWithCompositeAttributes struct {
	models.NoRangeThingWithCompositeAttributes
}

func (t NoRangeThingWithCompositeAttributesTable) name() string {
	if t.TableName != "" {
		return t.TableName
	}
	return fmt.Sprintf("%s-no-range-thing-with-composite-attributess", t.Prefix)
}

func (t NoRangeThingWithCompositeAttributesTable) create(ctx context.Context) error {
	if _, err := t.DynamoDBAPI.CreateTableWithContext(ctx, &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("date"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("name_branch"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("name_version"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("name_branch"),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			},
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("nameVersion"),
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("name_version"),
						KeyType:       aws.String(dynamodb.KeyTypeHash),
					},
					{
						AttributeName: aws.String("date"),
						KeyType:       aws.String(dynamodb.KeyTypeRange),
					},
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(t.ReadCapacityUnits),
					WriteCapacityUnits: aws.Int64(t.WriteCapacityUnits),
				},
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(t.ReadCapacityUnits),
			WriteCapacityUnits: aws.Int64(t.WriteCapacityUnits),
		},
		TableName: aws.String(t.name()),
	}); err != nil {
		return err
	}
	return nil
}

func (t NoRangeThingWithCompositeAttributesTable) saveNoRangeThingWithCompositeAttributes(ctx context.Context, m models.NoRangeThingWithCompositeAttributes) error {
	data, err := encodeNoRangeThingWithCompositeAttributes(m)
	if err != nil {
		return err
	}
	_, err = t.DynamoDBAPI.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(t.name()),
		Item:      data,
		ExpressionAttributeNames: map[string]*string{
			"#NAME_BRANCH": aws.String("name_branch"),
		},
		ConditionExpression: aws.String("attribute_not_exists(#NAME_BRANCH)"),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				return db.ErrNoRangeThingWithCompositeAttributesAlreadyExists{
					NameBranch: fmt.Sprintf("%s@%s", *m.Name, *m.Branch),
				}
			case dynamodb.ErrCodeResourceNotFoundException:
				return fmt.Errorf("table or index not found: %s", t.name())
			}
		}
		return err
	}
	return nil
}

func (t NoRangeThingWithCompositeAttributesTable) getNoRangeThingWithCompositeAttributes(ctx context.Context, name string, branch string) (*models.NoRangeThingWithCompositeAttributes, error) {
	key, err := dynamodbattribute.MarshalMap(ddbNoRangeThingWithCompositeAttributesPrimaryKey{
		NameBranch: fmt.Sprintf("%s@%s", name, branch),
	})
	if err != nil {
		return nil, err
	}
	res, err := t.DynamoDBAPI.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(t.name()),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				return nil, fmt.Errorf("table or index not found: %s", t.name())
			}
		}
		return nil, err
	}

	if len(res.Item) == 0 {
		return nil, db.ErrNoRangeThingWithCompositeAttributesNotFound{
			Name:   name,
			Branch: branch,
		}
	}

	var m models.NoRangeThingWithCompositeAttributes
	if err := decodeNoRangeThingWithCompositeAttributes(res.Item, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

func (t NoRangeThingWithCompositeAttributesTable) deleteNoRangeThingWithCompositeAttributes(ctx context.Context, name string, branch string) error {
	key, err := dynamodbattribute.MarshalMap(ddbNoRangeThingWithCompositeAttributesPrimaryKey{
		NameBranch: fmt.Sprintf("%s@%s", name, branch),
	})
	if err != nil {
		return err
	}
	_, err = t.DynamoDBAPI.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		Key:       key,
		TableName: aws.String(t.name()),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				return fmt.Errorf("table or index not found: %s", t.name())
			}
		}
		return err
	}

	return nil
}

func (t NoRangeThingWithCompositeAttributesTable) getNoRangeThingWithCompositeAttributessByNameVersionAndDate(ctx context.Context, input db.GetNoRangeThingWithCompositeAttributessByNameVersionAndDateInput, fn func(m *models.NoRangeThingWithCompositeAttributes, lastNoRangeThingWithCompositeAttributes bool) bool) error {
	if input.DateStartingAt != nil && input.StartingAfter != nil {
		return fmt.Errorf("Can specify only one of input.DateStartingAt or input.StartingAfter")
	}
	if input.Name == "" {
		return fmt.Errorf("Hash key input.Name cannot be empty")
	}
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(t.name()),
		IndexName: aws.String("nameVersion"),
		ExpressionAttributeNames: map[string]*string{
			"#NAME_VERSION": aws.String("name_version"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":nameVersion": &dynamodb.AttributeValue{
				S: aws.String(fmt.Sprintf("%s:%d", input.Name, input.Version)),
			},
		},
		ScanIndexForward: aws.Bool(!input.Descending),
		ConsistentRead:   aws.Bool(false),
	}
	if input.Limit != nil {
		queryInput.Limit = input.Limit
	}
	if input.DateStartingAt == nil {
		queryInput.KeyConditionExpression = aws.String("#NAME_VERSION = :nameVersion")
	} else {
		queryInput.ExpressionAttributeNames["#DATE"] = aws.String("date")
		queryInput.ExpressionAttributeValues[":date"] = &dynamodb.AttributeValue{
			S: aws.String(toDynamoTimeString(*input.DateStartingAt)),
		}
		if input.Descending {
			queryInput.KeyConditionExpression = aws.String("#NAME_VERSION = :nameVersion AND #DATE <= :date")
		} else {
			queryInput.KeyConditionExpression = aws.String("#NAME_VERSION = :nameVersion AND #DATE >= :date")
		}
	}
	if input.StartingAfter != nil {
		queryInput.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
			"date": &dynamodb.AttributeValue{
				S: aws.String(toDynamoTimeStringPtr(input.StartingAfter.Date)),
			},
			"name_version": &dynamodb.AttributeValue{
				S: aws.String(fmt.Sprintf("%s:%d", *input.StartingAfter.Name, input.StartingAfter.Version)),
			},
			"name_branch": &dynamodb.AttributeValue{
				S: aws.String(fmt.Sprintf("%s@%s", *input.StartingAfter.Name, *input.StartingAfter.Branch)),
			},
		}
	}

	var pageFnErr error
	pageFn := func(queryOutput *dynamodb.QueryOutput, lastPage bool) bool {
		if len(queryOutput.Items) == 0 {
			return false
		}
		items, err := decodeNoRangeThingWithCompositeAttributess(queryOutput.Items)
		if err != nil {
			pageFnErr = err
			return false
		}
		hasMore := true
		for i, item := range items {
			if lastPage == true {
				hasMore = i < len(items)-1
			}
			if !fn(&item, !hasMore) {
				return false
			}
		}
		return true
	}

	err := t.DynamoDBAPI.QueryPagesWithContext(ctx, queryInput, pageFn)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				return fmt.Errorf("table or index not found: %s", t.name())
			}
		}
		return err
	}
	if pageFnErr != nil {
		return pageFnErr
	}

	return nil
}

// encodeNoRangeThingWithCompositeAttributes encodes a NoRangeThingWithCompositeAttributes as a DynamoDB map of attribute values.
func encodeNoRangeThingWithCompositeAttributes(m models.NoRangeThingWithCompositeAttributes) (map[string]*dynamodb.AttributeValue, error) {
	val, err := dynamodbattribute.MarshalMap(ddbNoRangeThingWithCompositeAttributes{
		NoRangeThingWithCompositeAttributes: m,
	})
	if err != nil {
		return nil, err
	}
	// make sure composite attributes don't contain separator characters
	if strings.Contains(*m.Name, ":") {
		return nil, fmt.Errorf("name cannot contain ':': %s", *m.Name)
	}
	if strings.Contains(*m.Branch, "@") {
		return nil, fmt.Errorf("branch cannot contain '@': %s", *m.Branch)
	}
	if strings.Contains(*m.Name, "@") {
		return nil, fmt.Errorf("name cannot contain '@': %s", *m.Name)
	}
	// add in composite attributes
	primaryKey, err := dynamodbattribute.MarshalMap(ddbNoRangeThingWithCompositeAttributesPrimaryKey{
		NameBranch: fmt.Sprintf("%s@%s", *m.Name, *m.Branch),
	})
	if err != nil {
		return nil, err
	}
	for k, v := range primaryKey {
		val[k] = v
	}
	nameVersion, err := dynamodbattribute.MarshalMap(ddbNoRangeThingWithCompositeAttributesGSINameVersion{
		NameVersion: fmt.Sprintf("%s:%d", *m.Name, m.Version),
		Date:        *m.Date,
	})
	if err != nil {
		return nil, err
	}
	for k, v := range nameVersion {
		val[k] = v
	}
	return val, err
}

// decodeNoRangeThingWithCompositeAttributes translates a NoRangeThingWithCompositeAttributes stored in DynamoDB to a NoRangeThingWithCompositeAttributes struct.
func decodeNoRangeThingWithCompositeAttributes(m map[string]*dynamodb.AttributeValue, out *models.NoRangeThingWithCompositeAttributes) error {
	var ddbNoRangeThingWithCompositeAttributes ddbNoRangeThingWithCompositeAttributes
	if err := dynamodbattribute.UnmarshalMap(m, &ddbNoRangeThingWithCompositeAttributes); err != nil {
		return err
	}
	*out = ddbNoRangeThingWithCompositeAttributes.NoRangeThingWithCompositeAttributes
	return nil
}

// decodeNoRangeThingWithCompositeAttributess translates a list of NoRangeThingWithCompositeAttributess stored in DynamoDB to a slice of NoRangeThingWithCompositeAttributes structs.
func decodeNoRangeThingWithCompositeAttributess(ms []map[string]*dynamodb.AttributeValue) ([]models.NoRangeThingWithCompositeAttributes, error) {
	noRangeThingWithCompositeAttributess := make([]models.NoRangeThingWithCompositeAttributes, len(ms))
	for i, m := range ms {
		var noRangeThingWithCompositeAttributes models.NoRangeThingWithCompositeAttributes
		if err := decodeNoRangeThingWithCompositeAttributes(m, &noRangeThingWithCompositeAttributes); err != nil {
			return nil, err
		}
		noRangeThingWithCompositeAttributess[i] = noRangeThingWithCompositeAttributes
	}
	return noRangeThingWithCompositeAttributess, nil
}
