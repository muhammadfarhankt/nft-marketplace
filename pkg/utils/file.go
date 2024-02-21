package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func RandFileName(ext string) string {
	fileName := fmt.Sprintf("%s_%v", strings.ReplaceAll(uuid.NewString()[:6], "-", ""), time.Now().UnixMilli())
	if ext != "" {
		fileName = fmt.Sprintf("%s.%s", fileName, ext)
	}
	return fileName
}
