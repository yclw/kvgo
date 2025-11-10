package kvfile

import (
	"io"

	"github.com/yclw/kvgo"
)

const (
	FileType kvgo.Type = "file"
)

type FileValue struct {
	Value io.Reader
}

func (v *FileValue) Type() kvgo.Type {
	return FileType
}
