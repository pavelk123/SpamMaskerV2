package service

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type fileProducer struct {
	inputFile string
}

func (f fileProducer) produce() ([]string, error) {
	errPrefix := "fileProducer.produce:"

	f.inputFile = strings.TrimSuffix(f.inputFile, "\n")
	f.inputFile = strings.TrimSuffix(f.inputFile, "\r")
	file, err := os.OpenFile(f.inputFile, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("%s %w", errPrefix, err)
	}

	defer func() {
		if errDefer := file.Close(); errDefer != nil {
			err = fmt.Errorf("%s %w", errPrefix, errDefer)
		}
	}()

	var wr bytes.Buffer
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		if _, err = wr.WriteString(sc.Text()); err != nil {
			return nil, fmt.Errorf("%s %w", errPrefix, err)
		}
		wr.WriteString("\n")
	}

	return []string{strings.TrimSpace(wr.String())}, err
}

// NewFileProducer is constructor of fileProducer
func NewFileProducer(inputFile string) *fileProducer {
	return &fileProducer{inputFile: inputFile}
}
