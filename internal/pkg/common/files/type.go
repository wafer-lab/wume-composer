package files

import (
	"io"
	"strings"
)

var fileTypes = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
	"GIF87a":            "image/gif",
	"GIF89a":            "image/gif",
}

func getFileType(reader io.ReaderAt) string {
	buffer := make([]byte, 10)
	_, err := reader.ReadAt(buffer, 0)
	if err != nil {
		return ""
	}
	str := string(buffer)

	for magic, mime := range fileTypes {
		if strings.HasPrefix(str, magic) {
			return mime
		}
	}
	return ""
}

func isImage(mimeType string) bool {
	return strings.HasPrefix(mimeType, "image/")
}
