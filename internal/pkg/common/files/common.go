package files

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const MByte = 1 << 20

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func getFilenameFromUrl(path string) string {
	nameParts := strings.Split(path, "/")
	return nameParts[len(nameParts)-1]
}

func getFilename(path string) string {
	nameParts := strings.Split(path, string(os.PathSeparator))
	return nameParts[len(nameParts)-1]
}

func getExtension(path string) string {
	nameParts := strings.Split(path, ".")
	return nameParts[len(nameParts)-1]
}

func countMd5(reader io.Reader) (string, error) {
	h := md5.New()
	if _, err := io.Copy(h, reader); err != nil {
		return "", errors.New("Impossible to count md5! (" + err.Error() + ")")
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func countMd5FromFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", errors.New("Impossible to open file " + path + " to count md5! (" + err.Error() + ")")
	}

	hash, err := countMd5(file)
	if err := file.Close(); err != nil {
		return "", errors.New("Impossible to close file! (" + err.Error() + ")")
	}
	return hash, err
}
