package log

import (
	"fmt"
	"os"
)

func Write(msg string, localPath string) {
	fd, _ := os.OpenFile(localPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()

	content := fmt.Sprintf("%s\r\n", msg)
	buf := []byte(content)
	fd.Write(buf)
}
