package storage

import (
	"fmt"
	"os"
)

func WriteToFile(data []byte) error {
	f, err := os.OpenFile("./books.json",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("err open file: %s", err)
	}
	defer f.Close()
	if _, err := f.WriteString(string(data) + ",\n"); err != nil {
		return fmt.Errorf("err append file: %s", err)
	}
	return nil
}
