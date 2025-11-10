package string

import "github.com/yclw/kvgo"

const (
	StringType kvgo.Type = "string"
)

type StringValue struct {
	Value string
}

func (v *StringValue) Type() kvgo.Type {
	return StringType
}
