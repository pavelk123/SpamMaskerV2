package service

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

const defaultPath string = "./test/output.txt"

type FileProducer struct {
	inputFile string
}

func (f FileProducer) Produce() ([]string, error) {
	f.inputFile = strings.TrimSuffix(f.inputFile, "\n")
	f.inputFile = strings.TrimSuffix(f.inputFile, "\r")
	file, errFile := os.OpenFile(f.inputFile, os.O_RDONLY, 0666)
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

func NewFileProducer(inputFile string) *FileProducer {
	return &FileProducer{inputFile: inputFile}
}
