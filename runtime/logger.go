package runtime

import (
	"fmt"
	"os"
)

func AppendFile(path, content string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// write file
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	} else {
		// append file
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.WriteString("\n\n" + content)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteLog(hash string, content string) error {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		// create logs directory
		if err := os.Mkdir("logs", 0755); err != nil {
			return err
		}
	}
	return AppendFile(fmt.Sprintf("logs/%s.log", hash), content)
}
