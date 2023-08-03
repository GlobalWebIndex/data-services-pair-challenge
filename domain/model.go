package domain

import "encoding/json"

type (
	AudienceID    string
	AudienceName  string
	QuestionCode  string
	DatapointCode string
	SuffixCode    int
)

type Audience struct {
	ID         AudienceID   `json:"id,omitempty"`
	Name       AudienceName `json:"name,omitempty"`
	Expression Expression   `json:"expression,omitempty"`
}

// Option 1: write custom unmarshal function that allows you to decode
// 2 versions of Expression:
//
//	{"datapoint_codes": ["DP1", "DP2"]}
//	or
//	{"options": ["DP1", "DP2"]}
//
// The rest of the JSON is the same.
// See model_test.go for the test examples.
type Expression struct {
	And            []*Expression   `json:"and,omitempty"`
	Or             []*Expression   `json:"or,omitempty"`
	Not            bool            `json:"not,omitempty"`
	QuestionCode   QuestionCode    `json:"question_code,omitempty"`
	DatapointCodes []DatapointCode `json:"datapoint_codes,omitempty"`
	SuffixCodes    []SuffixCode    `json:"suffix_codes,omitempty"`
	MinCount       int             `json:"min_count,omitempty"`
}

// UnmarshalJSON customizes the JSON unmarshaling of the Expression struct.
// Option 1 solution:
func (e *Expression) UnmarshalJSON(data []byte) error {
	type ExpressionAlias Expression
	type ExpressionWithOptions struct {
		ExpressionAlias
		Options []DatapointCode `json:"options"`
	}

	var aliasedTarget ExpressionWithOptions
	if err := json.Unmarshal(data, &aliasedTarget); err != nil {
		return err
	}

	*e = Expression(aliasedTarget.ExpressionAlias)
	if len(aliasedTarget.DatapointCodes) == 0 && len(aliasedTarget.Options) != 0 {
		e.DatapointCodes = aliasedTarget.Options
	}

	return nil
}
