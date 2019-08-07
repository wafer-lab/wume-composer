package files

import (
	"mime/multipart"

	"wume-composer/internal/pkg/common/config"
	"wume-composer/internal/pkg/models"
)

var avatarsStorage = initStorage(config.Storage.Avatar.Dir, config.Storage.Avatar.Url)

func UploadAvatar(file multipart.File, header *multipart.FileHeader) (string, error) {
	if !isImage(header.Header.Get("Content-Type")) {
		return "", models.IncorrectDataError
	}

	if !isImage(getFileType(file)) {
		return "", models.IncorrectDataError
	}

	filename, err := avatarsStorage.SaveFileByExt(file, getExtension(header.Filename))
	if err != nil {
		return "", err
	}

	return avatarsStorage.url + filename, nil
}

func RemoveAvatar(url string) error {
	if url == "" {
		return nil
	}
	return avatarsStorage.Remove(getFilenameFromUrl(url))
}
