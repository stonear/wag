package dynamodb

import (
	"context"
	"fmt"
	"strings"

	"github.com/Clever/wag/samples/gen-go-db/models"
	"github.com/Clever/wag/samples/gen-go-db/server/db"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/go-openapi/strfmt"
)

var _ = strfmt.DateTime{}

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

// ddbTeacherSharingRuleGSIDistrictSchoolTeacherApp represents the district_school_teacher_app GSI.
type ddbTeacherSharingRuleGSIDistrictSchoolTeacherApp struct {
	District         string `dynamodbav:"district"`
	SchoolTeacherApp string `dynamodbav:"school_teacher_app"`
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
				AttributeName: aws.String("district"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("school_app"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("school_teacher_app"),
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
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("district_school_teacher_app"),
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("KEYS_ONLY"),
				},
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("district"),
						KeyType:       aws.String(dynamodb.KeyTypeHash),
					},
					{
						AttributeName: aws.String("school_teacher_app"),
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

func (t TeacherSharingRuleTable) getTeacherSharingRulesByTeacherAndSchoolApp(ctx context.Context, input db.GetTeacherSharingRulesByTeacherAndSchoolAppInput, fn func(m *models.TeacherSharingRule, lastTeacherSharingRule bool) bool) error {
	if input.StartingAt != nil && input.StartingAfter != nil {
		return fmt.Errorf("Can specify only one of StartingAt or StartingAfter")
	}
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
	if input.Limit != nil {
		queryInput.Limit = input.Limit
	}
	if input.StartingAt == nil {
		queryInput.KeyConditionExpression = aws.String("#TEACHER = :teacher")
	} else {
		queryInput.ExpressionAttributeNames["#SCHOOL_APP"] = aws.String("school_app")
		queryInput.ExpressionAttributeValues[":schoolApp"] = &dynamodb.AttributeValue{
			S: aws.String(fmt.Sprintf("%s_%s", input.StartingAt.School, input.StartingAt.App)),
		}
		if input.Descending {
			queryInput.KeyConditionExpression = aws.String("#TEACHER = :teacher AND #SCHOOL_APP <= :schoolApp")
		} else {
			queryInput.KeyConditionExpression = aws.String("#TEACHER = :teacher AND #SCHOOL_APP >= :schoolApp")
		}
	}
	if input.StartingAfter != nil {
		queryInput.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
			"school_app": &dynamodb.AttributeValue{
				S: aws.String(fmt.Sprintf("%s_%s", input.StartingAfter.School, input.StartingAfter.App)),
			},
			"teacher": &dynamodb.AttributeValue{
				S: aws.String(input.StartingAfter.Teacher),
			},
		}
	}

	var pageFnErr error
	pageFn := func(queryOutput *dynamodb.QueryOutput, lastPage bool) bool {
		if len(queryOutput.Items) == 0 {
			return false
		}
		items, err := decodeTeacherSharingRules(queryOutput.Items)
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
		return err
	}
	if pageFnErr != nil {
		return pageFnErr
	}

	return nil
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

func (t TeacherSharingRuleTable) getTeacherSharingRulesByDistrictAndSchoolTeacherApp(ctx context.Context, input db.GetTeacherSharingRulesByDistrictAndSchoolTeacherAppInput, fn func(m *models.TeacherSharingRule, lastTeacherSharingRule bool) bool) error {
	if input.StartingAt != nil && input.StartingAfter != nil {
		return fmt.Errorf("Can specify only one of input.StartingAt or input.StartingAfter")
	}
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(t.name()),
		IndexName: aws.String("district_school_teacher_app"),
		ExpressionAttributeNames: map[string]*string{
			"#DISTRICT": aws.String("district"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":district": &dynamodb.AttributeValue{
				S: aws.String(input.District),
			},
		},
		ScanIndexForward: aws.Bool(!input.Descending),
		ConsistentRead:   aws.Bool(false),
	}
	if input.Limit != nil {
		queryInput.Limit = input.Limit
	}
	if input.StartingAt == nil {
		queryInput.KeyConditionExpression = aws.String("#DISTRICT = :district")
	} else {
		queryInput.ExpressionAttributeNames["#SCHOOL_TEACHER_APP"] = aws.String("school_teacher_app")
		queryInput.ExpressionAttributeValues[":schoolTeacherApp"] = &dynamodb.AttributeValue{
			S: aws.String(fmt.Sprintf("%s_%s_%s", input.StartingAt.School, input.StartingAt.Teacher, input.StartingAt.App)),
		}
		if input.Descending {
			queryInput.KeyConditionExpression = aws.String("#DISTRICT = :district AND #SCHOOL_TEACHER_APP <= :schoolTeacherApp")
		} else {
			queryInput.KeyConditionExpression = aws.String("#DISTRICT = :district AND #SCHOOL_TEACHER_APP >= :schoolTeacherApp")
		}
	}
	if input.StartingAfter != nil {
		queryInput.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
			"school_teacher_app": &dynamodb.AttributeValue{
				S: aws.String(fmt.Sprintf("%s_%s_%s", input.StartingAfter.School, input.StartingAfter.Teacher, input.StartingAfter.App)),
			},
			"district": &dynamodb.AttributeValue{
				S: aws.String(input.StartingAfter.District),
			},
			"school_app": &dynamodb.AttributeValue{
				S: aws.String(fmt.Sprintf("%s_%s", input.StartingAfter.School, input.StartingAfter.App)),
			},
			"teacher": &dynamodb.AttributeValue{
				S: aws.String(input.StartingAfter.Teacher),
			},
		}
	}

	var pageFnErr error
	pageFn := func(queryOutput *dynamodb.QueryOutput, lastPage bool) bool {
		if len(queryOutput.Items) == 0 {
			return false
		}
		items, err := decodeTeacherSharingRules(queryOutput.Items)
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
		return err
	}
	if pageFnErr != nil {
		return pageFnErr
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
	// make sure composite attributes don't contain separator characters
	if strings.Contains(m.App, "_") {
		return nil, fmt.Errorf("app cannot contain '_': %s", m.App)
	}
	if strings.Contains(m.School, "_") {
		return nil, fmt.Errorf("school cannot contain '_': %s", m.School)
	}
	if strings.Contains(m.Teacher, "_") {
		return nil, fmt.Errorf("teacher cannot contain '_': %s", m.Teacher)
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
	districtSchoolTeacherApp, err := dynamodbattribute.MarshalMap(ddbTeacherSharingRuleGSIDistrictSchoolTeacherApp{
		District:         m.District,
		SchoolTeacherApp: fmt.Sprintf("%s_%s_%s", m.School, m.Teacher, m.App),
	})
	if err != nil {
		return nil, err
	}
	for k, v := range districtSchoolTeacherApp {
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
	// parse composite attributes from projected secondary indexes and fill
	// in model properties
	if v, ok := m["school_teacher_app"]; ok && v.S != nil {
		parts := strings.Split(*v.S, "_")
		if len(parts) != 3 {
			return fmt.Errorf("expected 3 parts: '%s'", *v.S)
		}
		out.School = parts[0]
		out.Teacher = parts[1]
		out.App = parts[2]
	}
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
