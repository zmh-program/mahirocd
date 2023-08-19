package utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"runtime"
)

func Md5Encode(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func Sha2Encode(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetCommandSeparator() string {
	system := runtime.GOOS

	if system == "windows" {
		if os.Getenv("PSModulePath") != "" || os.Getenv("PSVersionTable") != "" {
			return ";"
		}
		return "&"
	} else if system == "linux" {
		return "&&"
	} else {
		return ";"
	}
}
