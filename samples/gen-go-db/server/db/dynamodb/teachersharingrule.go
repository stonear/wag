package dynamodb

import (
	"context"
	"fmt"

	"github.com/Clever/wag/samples/gen-go-db/models"
	"github.com/Clever/wag/samples/gen-go-db/server/db"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// TeacherSharingRuleTable represents the user-configurable properties of the TeacherSharingRule table.
type TeacherSharingRuleTable struct {
	DynamoDBAPI        dynamodbiface.DynamoDBAPI
	Prefix             string
	TableName          string
	ReadCapacityUnits  int64
	WriteCapacityUnits int64
}

// ddbTeacherSharingRulePrimaryKey represents the primary key of a TeacherSharingRule in DynamoDB.
type ddbTeacherSharingRulePrimaryKey struct {
	Teacher   string `dynamodbav:"teacher"`
	SchoolApp string `dynamodbav:"school_app"`
}

// ddbTeacherSharingRule represents a TeacherSharingRule as stored in DynamoDB.
type ddbTeacherSharingRule struct {
	models.TeacherSharingRule
}

func (t TeacherSharingRuleTable) name() string {
	if t.TableName != "" {
		return t.TableName
	}
	return fmt.Sprintf("%s-teacher-sharing-rules", t.Prefix)
}

func (t TeacherSharingRuleTable) create(ctx context.Context) error {
	if _, err := t.DynamoDBAPI.CreateTableWithContext(ctx, &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("school_app"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("teacher"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("teacher"),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			},
			{
				AttributeName: aws.String("school_app"),
				KeyType:       aws.String(dynamodb.KeyTypeRange),
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

func (t TeacherSharingRuleTable) saveTeacherSharingRule(ctx context.Context, m models.TeacherSharingRule) error {
	data, err := encodeTeacherSharingRule(m)
	if err != nil {
		return err
	}
	_, err = t.DynamoDBAPI.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(t.name()),
		Item:      data,
	})
	return err
}

func (t TeacherSharingRuleTable) getTeacherSharingRule(ctx context.Context, teacher string, school string, app string) (*models.TeacherSharingRule, error) {
	key, err := dynamodbattribute.MarshalMap(ddbTeacherSharingRulePrimaryKey{
		Teacher:   teacher,
		SchoolApp: fmt.Sprintf("%s_%s", school, app),
	})
	if err != nil {
		return nil, err
	}
	res, err := t.DynamoDBAPI.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(t.name()),
	})
	if err != nil {
		return nil, err
	}

	if len(res.Item) == 0 {
		return nil, db.ErrTeacherSharingRuleNotFound{
			Teacher: teacher,
			School:  school,
			App:     app,
		}
	}

	var m models.TeacherSharingRule
	if err := decodeTeacherSharingRule(res.Item, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

func (t TeacherSharingRuleTable) getTeacherSharingRulesByTeacherAndSchoolApp(ctx context.Context, input db.GetTeacherSharingRulesByTeacherAndSchoolAppInput) ([]models.TeacherSharingRule, error) {
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(t.name()),
		ExpressionAttributeNames: map[string]*string{
			"#TEACHER": aws.String("teacher"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":teacher": &dynamodb.AttributeValue{
				S: aws.String(input.Teacher),
			},
		},
		ScanIndexForward: aws.Bool(!input.Descending),
		ConsistentRead:   aws.Bool(!input.DisableConsistentRead),
	}
	if input.StartingAt == nil {
		queryInput.KeyConditionExpression = aws.String("#TEACHER = :teacher")
	} else {
		queryInput.ExpressionAttributeNames["#SCHOOL_APP"] = aws.String("school_app")
		queryInput.ExpressionAttributeValues[":schoolApp"] = &dynamodb.AttributeValue{
			S: aws.String(fmt.Sprintf("%s_%s", input.StartingAt.School, input.StartingAt.App)),
		}
		queryInput.KeyConditionExpression = aws.String("#TEACHER = :teacher AND #SCHOOL_APP >= :schoolApp")
	}

	queryOutput, err := t.DynamoDBAPI.QueryWithContext(ctx, queryInput)
	if err != nil {
		return nil, err
	}
	if len(queryOutput.Items) == 0 {
		return []models.TeacherSharingRule{}, nil
	}

	return decodeTeacherSharingRules(queryOutput.Items)
}

func (t TeacherSharingRuleTable) deleteTeacherSharingRule(ctx context.Context, teacher string, school string, app string) error {
	key, err := dynamodbattribute.MarshalMap(ddbTeacherSharingRulePrimaryKey{
		Teacher:   teacher,
		SchoolApp: fmt.Sprintf("%s_%s", school, app),
	})
	if err != nil {
		return err
	}
	_, err = t.DynamoDBAPI.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		Key:       key,
		TableName: aws.String(t.name()),
	})
	if err != nil {
		return err
	}
	return nil
}

// encodeTeacherSharingRule encodes a TeacherSharingRule as a DynamoDB map of attribute values.
func encodeTeacherSharingRule(m models.TeacherSharingRule) (map[string]*dynamodb.AttributeValue, error) {
	val, err := dynamodbattribute.MarshalMap(ddbTeacherSharingRule{
		TeacherSharingRule: m,
	})
	if err != nil {
		return nil, err
	}
	// add in composite attributes
	primaryKey, err := dynamodbattribute.MarshalMap(ddbTeacherSharingRulePrimaryKey{
		Teacher:   m.Teacher,
		SchoolApp: fmt.Sprintf("%s_%s", m.School, m.App),
	})
	if err != nil {
		return nil, err
	}
	for k, v := range primaryKey {
		val[k] = v
	}
	return val, err
}

// decodeTeacherSharingRule translates a TeacherSharingRule stored in DynamoDB to a TeacherSharingRule struct.
func decodeTeacherSharingRule(m map[string]*dynamodb.AttributeValue, out *models.TeacherSharingRule) error {
	var ddbTeacherSharingRule ddbTeacherSharingRule
	if err := dynamodbattribute.UnmarshalMap(m, &ddbTeacherSharingRule); err != nil {
		return err
	}
	*out = ddbTeacherSharingRule.TeacherSharingRule
	return nil
}

// decodeTeacherSharingRules translates a list of TeacherSharingRules stored in DynamoDB to a slice of TeacherSharingRule structs.
func decodeTeacherSharingRules(ms []map[string]*dynamodb.AttributeValue) ([]models.TeacherSharingRule, error) {
	teacherSharingRules := make([]models.TeacherSharingRule, len(ms))
	for i, m := range ms {
		var teacherSharingRule models.TeacherSharingRule
		if err := decodeTeacherSharingRule(m, &teacherSharingRule); err != nil {
			return nil, err
		}
		teacherSharingRules[i] = teacherSharingRule
	}
	return teacherSharingRules, nil
}
