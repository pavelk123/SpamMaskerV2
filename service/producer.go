package service

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

type FileProducer struct {
}

func (f FileProducer) Produce(inputFile string) ([]string, error) {
	inputFile = strings.TrimSuffix(inputFile, "\n")
	inputFile = strings.TrimSuffix(inputFile, "\r")
	file, errFile := os.OpenFile(inputFile, os.O_RDONLY, 0666)
	if errFile != nil {
		return nil, errFile
	}
	defer file.Close()

	var wr bytes.Buffer
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		_, errScan := wr.WriteString(sc.Text())
		wr.WriteString("\n")
		if errScan != nil {
			return nil, errScan
		}
	}

	return []string{strings.TrimSpace(wr.String())}, nil
}

func NewFileProducer() *FileProducer {
	return &FileProducer{}
}
