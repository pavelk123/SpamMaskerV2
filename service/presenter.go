package service

import (
	"bufio"
	"fmt"
	"os"
)

const defaultPath string = "./test/output.txt"

type filePresenter struct {
	outputFile string
}

func (fp *filePresenter) present(data []string) error {
	
	file, err := os.Create(fp.outputFile)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}

	writer := bufio.NewWriter(file)
	defer func() {
		if errDefer := file.Close(); err != nil {
			err = fmt.Errorf("file.Close: %w", errDefer)
		}
	}()

	for _, str := range data {
		if _, err = writer.WriteString(str);err != nil {
			return fmt.Errorf("writer.WriteString: %w", err)
		}
	}

	if err = writer.Flush(); err != nil {
		return fmt.Errorf("writer.Flush: %w", err)
	}

	return err
}

// NewFilePresenter is constructor of filePresenter
// If path for output file is empty, then output file will be default
func NewFilePresenter(outputFile string) *filePresenter {
	if outputFile == "" {
		outputFile = defaultPath
	}

	return &filePresenter{outputFile: outputFile}
}
