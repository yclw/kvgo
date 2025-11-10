package file

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/yclw/kvgo"
)

type FileHandler struct {
	client *minio.Client
}

var _ kvgo.Handler = (*FileHandler)(nil)
var _ kvgo.Setter = (*FileHandler)(nil)
var _ kvgo.Getter = (*FileHandler)(nil)
var _ kvgo.Deleter = (*FileHandler)(nil)

func NewFileHandler(client *minio.Client) kvgo.Handler {
	return &FileHandler{client: client}
}

func (h *FileHandler) Set(namespace string, key string, val kvgo.Value, opts ...kvgo.Option) error {
	fileValue, ok := val.(*FileValue)
	if !ok {
		return fmt.Errorf("value is not a file")
	}
	option := GetOptions(opts...)
	ctx := option.Context
	fileOptions := GetFileOptions(opts...)
	_, err := h.client.PutObject(ctx, namespace, key, fileValue.Value, fileOptions.Size, minio.PutObjectOptions{ContentType: fileOptions.ContentType})
	if err != nil {
		return err
	}
	return nil
}

func (h *FileHandler) Get(namespace string, key string, opts ...kvgo.Option) (kvgo.Value, error) {
	option := GetOptions(opts...)
	ctx := option.Context
	object, err := h.client.GetObject(ctx, namespace, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return &FileValue{Value: object}, nil
}

func (h *FileHandler) Delete(namespace string, key string, opts ...kvgo.Option) error {
	option := GetOptions(opts...)
	ctx := option.Context
	return h.client.RemoveObject(ctx, namespace, key, minio.RemoveObjectOptions{})
}
