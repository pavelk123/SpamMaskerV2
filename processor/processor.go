package processor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Processor interface {
	Process(data []string) error
}

type StringProcessor struct{}

func (sp StringProcessor) Process(data []string) error {
	text := strings.Join(data, " ")
	fmt.Println(text)

	return nil
}

type FileProcessor struct {
	filePath string
}

func (fp *FileProcessor) Init(filepath string) {
	fp.filePath = filepath
}

func (fp *FileProcessor) Process(data []string) error {
	file, err := os.Create(fp.filePath)
	writer := bufio.NewWriter(file)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, str := range data {
		_, err = writer.WriteString(str)
		if err != nil {
			return err
		}
	}

	writer.Flush()
	return nil
}