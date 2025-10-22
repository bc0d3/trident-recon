package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

// GenerateID generates a unique ID for a session
func GenerateID(tool, commandName, target string) string {
	data := fmt.Sprintf("%s_%s_%s_%d", tool, commandName, target, time.Now().UnixNano())
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])[:12]
}
