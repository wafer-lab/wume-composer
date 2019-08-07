package files

import (
	"os"
	"path/filepath"

	"wume-composer/internal/pkg/common/config"
)

var trashStorage = initStorage(config.Storage.Trash.Dir, config.Storage.Trash.Url)

func moveInTrash(path string) error {
	if !exists(path) {
		return nil
	}

	hash, err := countMd5FromFile(path)
	if err != nil {
		return err
	}

	trashFilename := filepath.Join(trashStorage.path, hash+"."+getExtension(path))
	if exists(trashFilename) {
		return os.Remove(path)
	} else {
		return os.Rename(path, trashFilename)
	}
}
