package service

import (
	"os"
	"bufio"
)

type FilePresenter struct{
	outputFile string
}

func (fp *FilePresenter) Present(data []string) error {
	file, err := os.Create(fp.outputFile)
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

func NewFilePresenter(outputFile string)*FilePresenter{
	return &FilePresenter{outputFile: outputFile}
}

