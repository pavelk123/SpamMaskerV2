package service

import (
	"bufio"
	"os"
)

type FilePresenter struct {
	outputFile string
}

func (fp *FilePresenter) Present(data []string) error {
	file, err := os.Create(fp.outputFile)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	defer file.Close()

	for _, str := range data {
		_, err = writer.WriteString(str)
		if err != nil {
			return err
		}
	}

	if err = writer.Flush();err != nil { 
		return err
	}

	return nil
}

func NewFilePresenter(outputFile string) *FilePresenter {
	if outputFile == "" {
		outputFile = defaultPath
	}

	return &FilePresenter{outputFile: outputFile}
}
