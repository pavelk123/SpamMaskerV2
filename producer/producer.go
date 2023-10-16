package producer

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

type Producer interface {
	Produce(data string) ([]string, error)
}

type StringProducer struct{}

func (sp StringProducer) Produce(data string) ([]string, error) {
	return []string{strings.TrimSpace(data)}, nil
}

type FileProducer struct {
	filePath string
}

func (fp *FileProducer) Produce(data string) ([]string, error) {
	data = strings.TrimSuffix(data, "\n")
	data = strings.TrimSuffix(data, "\r")
	file, errFile := os.OpenFile(data, os.O_RDONLY, 0666)
	if errFile != nil {
		return nil, errFile
	}
	defer file.Close()

	fp.filePath = data

	var wr bytes.Buffer
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		wr.WriteString(sc.Text())
		wr.WriteString("\n")
	}

	return []string{strings.TrimSpace(wr.String())}, nil
}

func (fp *FileProducer) GetFilePath() string {
	return fp.filePath
}
