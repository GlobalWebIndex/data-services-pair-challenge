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
	ID         AudienceID
	Name       AudienceName
	Expression Expression
}

type Expression struct {
	And            []*Expression
	Or             []*Expression
	Not            bool
	QuestionCode   QuestionCode
	DatapointCodes []DatapointCode
	SuffixCodes    []SuffixCode
	MinCount       int
}

// MarshalJSON customizes the JSON marshaling of the Audience struct.
func (a Audience) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID         string     `json:"id"`
		Name       string     `json:"name"`
		Expression Expression `json:"expression"`
	}{
		ID:         string(a.ID),
		Name:       string(a.Name),
		Expression: a.Expression,
	})
}

// UnmarshalJSON customizes the JSON unmarshaling of the Audience struct.
func (a *Audience) UnmarshalJSON(data []byte) error {
	aux := &struct {
		ID         string     `json:"id"`
		Name       string     `json:"name"`
		Expression Expression `json:"expression"`
	}{
		ID:         "",
		Name:       "",
		Expression: Expression{},
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	a.ID = AudienceID(aux.ID)
	a.Name = AudienceName(aux.Name)
	a.Expression = aux.Expression
	return nil
}

// MarshalJSON customizes the JSON marshaling of the Expression struct.
func (e Expression) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		And            []*ExpressionAlias `json:"and,omitempty"`
		Or             []*ExpressionAlias `json:"or,omitempty"`
		Not            bool               `json:"not,omitempty"`
		QuestionCode   string             `json:"question_code,omitempty"`
		DatapointCodes []string           `json:"datapoint_codes,omitempty"`
		SuffixCodes    []int              `json:"suffix_codes,omitempty"`
		MinCount       int                `json:"min_count,omitempty"`
	}{
		And:            toExpressionAlias(e.And),
		Or:             toExpressionAlias(e.Or),
		Not:            e.Not,
		QuestionCode:   string(e.QuestionCode),
		DatapointCodes: toStringSlice(e.DatapointCodes),
		SuffixCodes:    toIntSlice(e.SuffixCodes),
		MinCount:       e.MinCount,
	})
}

// UnmarshalJSON customizes the JSON unmarshaling of the Expression struct.
func (e *Expression) UnmarshalJSON(data []byte) error {
	aux := struct {
		And            []*ExpressionAlias `json:"and,omitempty"`
		Or             []*ExpressionAlias `json:"or,omitempty"`
		Not            bool               `json:"not"`
		QuestionCode   string             `json:"question_code"`
		DatapointCodes []string           `json:"datapoint_codes"`
		SuffixCodes    []int              `json:"suffix_codes"`
		MinCount       int                `json:"min_count"`
	}{
		And:            []*ExpressionAlias{},
		Or:             []*ExpressionAlias{},
		Not:            false,
		QuestionCode:   "",
		DatapointCodes: []string{},
		SuffixCodes:    []int{},
		MinCount:       0,
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	e.And = expressionAliasToExpression(aux.And)
	e.Or = expressionAliasToExpression(aux.Or)
	e.Not = aux.Not
	e.QuestionCode = QuestionCode(aux.QuestionCode)
	e.DatapointCodes = toDatapointCodeSlice(aux.DatapointCodes)
	e.SuffixCodes = toSuffixCodeSlice(aux.SuffixCodes)
	e.MinCount = aux.MinCount
	return nil
}

// ExpressionAlias is an alias for Expression.
type ExpressionAlias = Expression

// ToAlias converts an Expression slice to Alias slice.
func toExpressionAlias(expressions []*Expression) []*ExpressionAlias {
	if expressions == nil {
		return nil
	}
	var aliases []*ExpressionAlias
	return append(aliases, expressions...)
}

// ToExpression converts an Alias slice to Expression slice.
func expressionAliasToExpression(aliases []*ExpressionAlias) []*Expression {
	if aliases == nil {
		return nil
	}
	var expressions []*Expression
	return append(expressions, aliases...)
}

// toStringSlice converts a slice of string to a pointer to the slice.
func toStringSlice(s []DatapointCode) []string {
	if len(s) == 0 {
		return nil
	}
	stringSlice := make([]string, len(s))
	for i, v := range s {
		stringSlice[i] = string(v)
	}
	return stringSlice
}

// toDatapointCodeSlice converts a slice of string to a pointer to the slice.
func toDatapointCodeSlice(s []string) []DatapointCode {
	if len(s) == 0 {
		return nil
	}
	datapointSlice := make([]DatapointCode, len(s))
	for i, v := range s {
		datapointSlice[i] = DatapointCode(v)
	}
	return datapointSlice
}

// toIntSlice converts a slice of int to a pointer to the slice.
func toIntSlice(s []SuffixCode) []int {
	if len(s) == 0 {
		return nil
	}
	intSlice := make([]int, len(s))
	for i, v := range s {
		intSlice[i] = int(v)
	}
	return intSlice
}

// toSuffixCodeSlice converts a slice of int to a pointer to the slice.
func toSuffixCodeSlice(s []int) []SuffixCode {
	if len(s) == 0 {
		return nil
	}
	suffixSlice := make([]SuffixCode, len(s))
	for i, v := range s {
		suffixSlice[i] = SuffixCode(v)
	}
	return suffixSlice
}
