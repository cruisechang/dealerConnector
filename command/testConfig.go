package command

import (
	"bytes"
	"os"
	"path/filepath"
)

func getConfigFilePosition(fileName string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
	}

	var buf bytes.Buffer
	buf.WriteString(dir)
	buf.WriteString("/")
	buf.WriteString(fileName)

	return buf.String()
}
