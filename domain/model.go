package domain

type (
	AudienceID    string
	AudienceName  string
	QuestionCode  string
	DatapointCode string
	SuffixCode    int
)

type Audience struct {
	ID         AudienceID
	Name       AudienceName
	Expression Expression
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
	And            []*Expression
	Or             []*Expression
	Not            bool
	QuestionCode   QuestionCode
	DatapointCodes []DatapointCode
	SuffixCodes    []SuffixCode
	MinCount       int
}
