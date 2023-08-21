package utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"runtime"
	"time"
)

func Md5Encode(str string) string {
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

func SetTimeout(f func(), timeout int) {
	go func() {
		<-time.After(time.Duration(timeout) * time.Millisecond)
		f()
	}()
}

func SetInterval(f func(), interval int) {
	go func() {
		for {
			<-time.After(time.Duration(interval) * time.Millisecond)
			f()
		}
	}()
}

func SetTimeoutSync(f func(), timeout int) {
	<-time.After(time.Duration(timeout) * time.Millisecond)
	f()
}
