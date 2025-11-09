package kvgo

type Type string

type Value interface {
	Type() Type
}
