package domain_test

import (
	"encoding/json"
	"testing"

	"github.com/GlobalWebIndex/data-services-pair-challenge/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestAudienceEncodingAndDecoding(t *testing.T) {
	// Create an example Audience instance.
	originalAudience := domain.Audience{
		ID:   "1",
		Name: "Test Audience",
		Expression: domain.Expression{
			Not:          true,
			QuestionCode: "Q1",
			DatapointCodes: []domain.DatapointCode{
				"DP1",
				"DP2",
			},
			SuffixCodes: []domain.SuffixCode{
				100,
				200,
			},
			MinCount: 10,
		},
	}

	// Marshal Audience to JSON.
	jsonData, err := json.Marshal(originalAudience)
	assert.NoError(t, err, "Error while marshaling Audience")

	// Unmarshal JSON back to Audience.
	var decodedAudience domain.Audience
	err = json.Unmarshal(jsonData, &decodedAudience)
	assert.NoError(t, err, "Error while unmarshaling Audience")

	// Compare the original and decoded Audience.
	assert.Equal(t, originalAudience, decodedAudience, "Decoded Audience does not match the original")
}

func TestExpressionEncodingAndDecoding(t *testing.T) {
	// Create an example Expression instance.
	originalExpression := domain.Expression{
		Not:          true,
		QuestionCode: "Q1",
		DatapointCodes: []domain.DatapointCode{
			"DP1",
			"DP2",
		},
		SuffixCodes: []domain.SuffixCode{
			100,
			200,
		},
		MinCount: 10,
	}

	// Marshal Expression to JSON.
	jsonData, err := json.Marshal(originalExpression)
	assert.NoError(t, err, "Error while marshaling Expression")

	// Unmarshal JSON back to Expression.
	var decodedExpression domain.Expression
	err = json.Unmarshal(jsonData, &decodedExpression)
	assert.NoError(t, err, "Error while unmarshaling Expression")

	// Compare the original and decoded Expression.
	assert.Equal(t, originalExpression.And, decodedExpression.And, "Decoded 'And' field does not match the original")
	assert.Equal(t, originalExpression.Or, decodedExpression.Or, "Decoded 'Or' field does not match the original")
	assert.Equal(t, originalExpression.Not, decodedExpression.Not, "Decoded 'Not' field does not match the original")
	assert.Equal(t, originalExpression.QuestionCode, decodedExpression.QuestionCode, "Decoded 'QuestionCode' field does not match the original")
	assert.Equal(t, originalExpression.DatapointCodes, decodedExpression.DatapointCodes, "Decoded 'DatapointCodes' field does not match the original")
	assert.Equal(t, originalExpression.SuffixCodes, decodedExpression.SuffixCodes, "Decoded 'SuffixCodes' field does not match the original")
	assert.Equal(t, originalExpression.MinCount, decodedExpression.MinCount, "Decoded 'MinCount' field does not match the original")
}

type AudienceSuite struct {
	suite.Suite
}

func TestAudienceSuite(t *testing.T) {
	suite.Run(t, new(AudienceSuite))
}

func (suite *AudienceSuite) TestAudienceJSONToStruct() {
	jsonData := []byte(`
	{
		"id": "1",
		"name": "Test Audience",
		"expression": {
			"and": [
				{
					"question_code": "Q1",
					"datapoint_codes": ["DP1", "DP2"],
					"suffix_codes": [100, 200],
					"min_count": 10,
					"not": true
				},
				{
					"question_code": "Q2",
					"datapoint_codes": ["DP3", "DP4"],
					"suffix_codes": [300, 400],
					"min_count": 20
				}
			]
		}
	}`)

	var audience domain.Audience
	err := json.Unmarshal(jsonData, &audience)
	assert.NoError(suite.T(), err, "Error while unmarshaling JSON to Audience")

	expectedAudience := domain.Audience{
		ID:   "1",
		Name: "Test Audience",
		Expression: domain.Expression{
			And: []*domain.Expression{
				{
					QuestionCode:   "Q1",
					DatapointCodes: []domain.DatapointCode{"DP1", "DP2"},
					SuffixCodes:    []domain.SuffixCode{100, 200},
					MinCount:       10,
					Not:            true,
				},
				{
					QuestionCode:   "Q2",
					DatapointCodes: []domain.DatapointCode{"DP3", "DP4"},
					SuffixCodes:    []domain.SuffixCode{300, 400},
					MinCount:       20,
				},
			},
		},
	}

	suite.Equal(expectedAudience, audience, "Decoded Audience does not match the expected")
}

func (suite *AudienceSuite) TestAudienceStructToJSON() {
	audience := domain.Audience{
		ID:   "1",
		Name: "Test Audience",
		Expression: domain.Expression{
			Or: []*domain.Expression{
				{
					QuestionCode:   "Q1",
					DatapointCodes: []domain.DatapointCode{"DP1", "DP2"},
					SuffixCodes:    []domain.SuffixCode{100, 200},
					MinCount:       10,
					Not:            true,
				},
				{
					QuestionCode:   "Q2",
					DatapointCodes: []domain.DatapointCode{"DP3", "DP4"},
					SuffixCodes:    []domain.SuffixCode{300, 400},
					MinCount:       20,
				},
			},
		},
	}

	expectedJSON := []byte(`
	{
		"id": "1",
		"name": "Test Audience",
		"expression": {
			"or": [
				{
					"question_code": "Q1",
					"datapoint_codes": ["DP1", "DP2"],
					"suffix_codes": [100, 200],
					"min_count": 10,
					"not": true
				},
				{
					"question_code": "Q2",
					"datapoint_codes": ["DP3", "DP4"],
					"suffix_codes": [300, 400],
					"min_count": 20
				}
			]
		}
	}`)

	jsonData, err := json.Marshal(audience)
	suite.NoError(err, "Error while marshaling Audience to JSON")
	suite.JSONEq(string(expectedJSON), string(jsonData), "Encoded JSON does not match the expected")
}
