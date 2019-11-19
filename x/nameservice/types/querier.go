package types

import (
	"fmt"
	"strings"
)

type QueryResResolve struct {
	Value string `json:"value"`
}

func (resolve QueryResResolve) String() string {
	return resolve.Value
}

var _ fmt.Stringer = new(QueryResResolve)

type QueryResNames []string

func (names QueryResNames) String() string {
	return strings.Join(names[:], "\n")
}

var _ fmt.Stringer = new(QueryResNames)
