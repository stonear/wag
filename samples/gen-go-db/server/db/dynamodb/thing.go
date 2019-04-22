package dynamodb

import (
	"context"
	"fmt"

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

// ThingTable represents the user-configurable properties of the Thing table.
type ThingTable struct {
	DynamoDBAPI        dynamodbiface.DynamoDBAPI
	Prefix             string
	TableName          string
	ReadCapacityUnits  int64
	WriteCapacityUnits int64
}

// ddbThingPrimaryKey represents the primary key of a Thing in DynamoDB.
type ddbThingPrimaryKey struct {
	Name    string `dynamodbav:"name"`
	Version int64  `dynamodbav:"version"`
}

// ddbThingGSIThingID represents the thingID GSI.
type ddbThingGSIThingID struct {
	ID string `dynamodbav:"id"`
}

// ddbThingGSINameCreatedAt represents the name-createdAt GSI.
type ddbThingGSINameCreatedAt struct {
	Name      string          `dynamodbav:"name"`
	CreatedAt strfmt.DateTime `dynamodbav:"createdAt"`
}

// ddbThing represents a Thing as stored in DynamoDB.
type ddbThing struct {
	models.Thing
}

func (t ThingTable) name() string {
	if t.TableName != "" {
		return t.TableName
	}
	return fmt.Sprintf("%s-things", t.Prefix)
}

func (t ThingTable) create(ctx context.Context) error {
	if _, err := t.DynamoDBAPI.CreateTableWithContext(ctx, &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("createdAt"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("name"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("version"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("name"),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			},
			{
				AttributeName: aws.String("version"),
				KeyType:       aws.String(dynamodb.KeyTypeRange),
			},
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("thingID"),
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("id"),
						KeyType:       aws.String(dynamodb.KeyTypeHash),
					},
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(t.ReadCapacityUnits),
					WriteCapacityUnits: aws.Int64(t.WriteCapacityUnits),
				},
			},
			{
				IndexName: aws.String("name-createdAt"),
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("name"),
						KeyType:       aws.String(dynamodb.KeyTypeHash),
					},
					{
						AttributeName: aws.String("createdAt"),
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

func (t ThingTable) saveThing(ctx context.Context, m models.Thing) error {
	data, err := encodeThing(m)
	if err != nil {
		return err
	}
	_, err = t.DynamoDBAPI.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(t.name()),
		Item:      data,
		ExpressionAttributeNames: map[string]*string{
			"#NAME":    aws.String("name"),
			"#VERSION": aws.String("version"),
		},
		ConditionExpression: aws.String("attribute_not_exists(#NAME) AND attribute_not_exists(#VERSION)"),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				return db.ErrThingAlreadyExists{
					Name:    m.Name,
					Version: m.Version,
				}
			case dynamodb.ErrCodeResourceNotFoundException:
				return fmt.Errorf("table or index not found: %s", t.name())
			}
		}
		return err
	}
	return nil
}

func (t ThingTable) getThing(ctx context.Context, name string, version int64) (*models.Thing, error) {
	key, err := dynamodbattribute.MarshalMap(ddbThingPrimaryKey{
		Name:    name,
		Version: version,
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
		return nil, db.ErrThingNotFound{
			Name:    name,
			Version: version,
		}
	}

	var m models.Thing
	if err := decodeThing(res.Item, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

func (t ThingTable) scanThings(ctx context.Context, input db.ScanThingsInput, fn func(m *models.Thing, lastThing bool) bool) error {
	scanInput := &dynamodb.ScanInput{
		TableName:      aws.String(t.name()),
		ConsistentRead: aws.Bool(!input.DisableConsistentRead),
	}
	if input.StartingAfter != nil {
		exclusiveStartKey, err := dynamodbattribute.MarshalMap(input.StartingAfter)
		if err != nil {
			return fmt.Errorf("error encoding exclusive start key for scan: %s", err.Error())
		}
		// must provide only the fields constituting the index
		scanInput.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
			"name":    exclusiveStartKey["name"],
			"version": exclusiveStartKey["version"],
		}
	}
	var innerErr error
	err := t.DynamoDBAPI.ScanPagesWithContext(ctx, scanInput, func(out *dynamodb.ScanOutput, lastPage bool) bool {
		ms, err := decodeThings(out.Items)
		if err != nil {
			innerErr = fmt.Errorf("error decoding %s", err.Error())
			return false
		}
		for i := range ms {
			if input.Limiter != nil {
				if err := input.Limiter.Wait(ctx); err != nil {
					innerErr = err
					return false
				}
			}
			lastModel := lastPage && i == len(ms)-1
			if continuee := fn(&ms[i], lastModel); !continuee {
				return false
			}
		}
		return true
	})
	if innerErr != nil {
		return innerErr
	}
	return err
}

func (t ThingTable) getThingsByNameAndVersion(ctx context.Context, input db.GetThingsByNameAndVersionInput, fn func(m *models.Thing, lastThing bool) bool) error {
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(t.name()),
		ExpressionAttributeNames: map[string]*string{
			"#NAME":    aws.String("name"),
			"#VERSION": aws.String("version"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name": &dynamodb.AttributeValue{
				S: aws.String(input.StartingAt.Name),
			},
			":version": &dynamodb.AttributeValue{
				N: aws.String(fmt.Sprintf("%d", input.StartingAt.Version)),
			},
		},
		ScanIndexForward: aws.Bool(!input.Descending),
		ConsistentRead:   aws.Bool(!input.DisableConsistentRead),
		Limit:            input.Limit,
	}
	if input.Exclusive {
		queryInput.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
			"version": &dynamodb.AttributeValue{
				N: aws.String(fmt.Sprintf("%d", input.StartingAt.Version)),
			},
			"name": &dynamodb.AttributeValue{
				S: aws.String(input.StartingAt.Name),
			},
		}
	}
	if input.Descending {
		queryInput.KeyConditionExpression = aws.String("#NAME = :name AND #VERSION <= :version")
	} else {
		queryInput.KeyConditionExpression = aws.String("#NAME = :name AND #VERSION >= :version")
	}

	queryOutput, err := t.DynamoDBAPI.QueryWithContext(ctx, queryInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				return fmt.Errorf("table or index not found: %s", t.name())
			}
		}
		return err
	}
	if len(queryOutput.Items) == 0 {
		return nil
	}

	items, err := decodeThings(queryOutput.Items)
	if err != nil {
		return err
	}

	for i, item := range items {
		hasMore := false
		if len(queryOutput.LastEvaluatedKey) > 0 {
			hasMore = true
		} else {
			hasMore = i < len(items)-1
		}
		if !fn(&item, !hasMore) {
			break
		}
	}

	return nil
}

func (t ThingTable) deleteThing(ctx context.Context, name string, version int64) error {
	key, err := dynamodbattribute.MarshalMap(ddbThingPrimaryKey{
		Name:    name,
		Version: version,
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

func (t ThingTable) getThingByID(ctx context.Context, id string) (*models.Thing, error) {
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(t.name()),
		IndexName: aws.String("thingID"),
		ExpressionAttributeNames: map[string]*string{
			"#ID": aws.String("id"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": &dynamodb.AttributeValue{
				S: aws.String(id),
			},
		},
		KeyConditionExpression: aws.String("#ID = :id"),
	}

	queryOutput, err := t.DynamoDBAPI.QueryWithContext(ctx, queryInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				return nil, fmt.Errorf("table or index not found: %s", t.name())
			}
		}
		return nil, err
	}
	if len(queryOutput.Items) == 0 {
		return nil, db.ErrThingByIDNotFound{ID: id}
	}

	var thing models.Thing
	if err := decodeThing(queryOutput.Items[0], &thing); err != nil {
		return nil, err
	}
	return &thing, nil
}

func (t ThingTable) getThingsByNameAndCreatedAt(ctx context.Context, input db.GetThingsByNameAndCreatedAtInput, fn func(m *models.Thing, lastThing bool) bool) error {
	if input.StartingAt == nil {
		return fmt.Errorf("StartingAt cannot be nil")
	}
	if input.Limit == nil {
		return fmt.Errorf("Limit cannot be nil")
	}
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(t.name()),
		IndexName: aws.String("name-createdAt"),
		ExpressionAttributeNames: map[string]*string{
			"#NAME":      aws.String("name"),
			"#CREATEDAT": aws.String("createdAt"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name": &dynamodb.AttributeValue{
				S: aws.String(input.StartingAt.Name),
			},
			":createdAt": &dynamodb.AttributeValue{
				S: aws.String(toDynamoTimeString(input.StartingAt.CreatedAt)),
			},
		},
		ScanIndexForward: aws.Bool(!input.Descending),
		ConsistentRead:   aws.Bool(false),
		Limit:            input.Limit,
	}
	if input.Exclusive {
		queryInput.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
			"createdAt": &dynamodb.AttributeValue{
				S: aws.String(toDynamoTimeString(input.StartingAt.CreatedAt)),
			},
			"name": &dynamodb.AttributeValue{
				S: aws.String(input.StartingAt.Name),
			},

			"version": &dynamodb.AttributeValue{
				N: aws.String(fmt.Sprintf("%d", input.StartingAt.Version)),
			},
		}
	}
	if input.Descending {
		queryInput.KeyConditionExpression = aws.String("#NAME = :name AND #CREATEDAT <= :createdAt")
	} else {
		queryInput.KeyConditionExpression = aws.String("#NAME = :name AND #CREATEDAT >= :createdAt")
	}

	queryOutput, err := t.DynamoDBAPI.QueryWithContext(ctx, queryInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				return fmt.Errorf("table or index not found: %s", t.name())
			}
		}
		return err
	}
	if len(queryOutput.Items) == 0 {
		return nil
	}

	items, err := decodeThings(queryOutput.Items)
	if err != nil {
		return err
	}

	for i, item := range items {
		hasMore := false
		if len(queryOutput.LastEvaluatedKey) > 0 {
			hasMore = true
		} else {
			hasMore = i < len(items)-1
		}
		if !fn(&item, !hasMore) {
			break
		}
	}

	return nil
}

// encodeThing encodes a Thing as a DynamoDB map of attribute values.
func encodeThing(m models.Thing) (map[string]*dynamodb.AttributeValue, error) {
	return dynamodbattribute.MarshalMap(ddbThing{
		Thing: m,
	})
}

// decodeThing translates a Thing stored in DynamoDB to a Thing struct.
func decodeThing(m map[string]*dynamodb.AttributeValue, out *models.Thing) error {
	var ddbThing ddbThing
	if err := dynamodbattribute.UnmarshalMap(m, &ddbThing); err != nil {
		return err
	}
	*out = ddbThing.Thing
	return nil
}

// decodeThings translates a list of Things stored in DynamoDB to a slice of Thing structs.
func decodeThings(ms []map[string]*dynamodb.AttributeValue) ([]models.Thing, error) {
	things := make([]models.Thing, len(ms))
	for i, m := range ms {
		var thing models.Thing
		if err := decodeThing(m, &thing); err != nil {
			return nil, err
		}
		things[i] = thing
	}
	return things, nil
}
