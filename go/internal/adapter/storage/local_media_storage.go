package storage

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/yohagos/multi-content-management/internal/core/port"
)

type localMediaStorage struct {
	basePath string
}

func NewLocalMediaStorage(basePath string) port.MediaStorage {
	return &localMediaStorage{
		basePath: basePath,
	}
}

func (s *localMediaStorage) Save(
	ctx context.Context,
	file io.Reader,
	fileName string,
) (string, int64, error) {
	if err := ctx.Err(); err != nil {
		return "", 0, err
	}

	if err := os.MkdirAll(s.basePath, 0755); err != nil {
		return "", 0, err
	}

	ext := filepath.Ext(filepath.Base(fileName))

	dst, err := os.CreateTemp(s.basePath, "media-*"+ext)
	if err != nil {
		return "", 0, err
	}

	defer dst.Close()

	filePath := filepath.Base(dst.Name())
	size, err := io.Copy(dst, file)
	if err != nil {
		_ = os.Remove(dst.Name())
		return "", 0, err
	}

	if err := ctx.Err(); err != nil {
		_ = os.Remove(dst.Name())
		return "", 0, err
	}

	if err := dst.Close(); err != nil {
		_ = os.Remove(dst.Name())
		return "", 0, err
	}

	return filePath, size, nil
}

func (s *localMediaStorage) Open(
	ctx context.Context,
	filePath string,
) (io.ReadCloser, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if !validMediaPath(filePath) {
		return nil, errors.New("invalid media path")
	}

	return os.Open(filepath.Join(s.basePath, filePath))
}

func (s *localMediaStorage) Delete(
	ctx context.Context,
	filePath string,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if filePath == "" || filepath.Base(filePath) != filePath {
		return errors.New("invalid media file path")
	}

	err := os.Remove(filepath.Join(s.basePath, filePath))
	if err != nil {
		return err
	}
	return nil
}

func validMediaPath(filePath string) bool {
	return filePath != "" &&
		filePath != "." &&
		filePath != ".." &&
		filepath.Base(filePath) == filePath
}

type contextReader struct {
	ctx context.Context
	r   io.Reader
}

func (r *contextReader) Read(p []byte) (int, error) {
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}

	return r.r.Read(p)
}
