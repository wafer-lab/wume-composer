package files

import (
	"io"
	"os"
	"path/filepath"

	"wume-composer/internal/pkg/common/logger"
	"wume-composer/internal/pkg/common/random"
)

type FileStorage struct {
	path string
	url  string
}

func (storage *FileStorage) SaveFileByExt(data io.Reader, ext string) (string, error) {
	for {
		filename := random.String(32) + "." + ext
		path := filepath.Join(storage.path, filename)
		if !exists(path) {
			return filename, storage.SaveFile(data, filename)
		}
	}
}

func (storage *FileStorage) SaveFile(data io.Reader, filename string) error {
	path := filepath.Join(storage.path, filename)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, data)
	if err := file.Close(); err != nil {
		logger.Error("Impossible to close new file: " + err.Error())
	}
	return err
}

func (storage *FileStorage) Remove(filename string) error {
	return moveInTrash(filepath.Join(storage.path, filename))
}

func newFileStorage(path, url string) *FileStorage {
	return &FileStorage{
		path: path,
		url:  url,
	}
}

func initStorage(dir string, url string) *FileStorage {
	var path, err = filepath.Abs(dir)
	if err != nil {
		logger.Error("Incorrect path of dir " + dir + "! (" + err.Error() + ")")
	}
	if !exists(path) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logger.Error("Impossible to create dir " + path + "! (" + err.Error() + ")")
		}
	}
	return newFileStorage(path, url)
}
