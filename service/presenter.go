package service

import (
	"bufio"
	"fmt"
	"os"
)

const defaultPath string = "./test/output.txt"

type FilePresenter struct {
	outputFile string
}

func (fp *FilePresenter) present(data []string) error {
	file, err := os.Create(fp.outputFile)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}

	writer := bufio.NewWriter(file)

	for _, str := range data {
		if _, err = writer.WriteString(str); err != nil {
			return fmt.Errorf("writer.WriteString: %w", err)
		}
	}

	if err = writer.Flush(); err != nil {
		return fmt.Errorf("writer.Flush: %w", err)
	}

	if err = file.Close(); err != nil {
		return fmt.Errorf("file.Close: %w", err)
	}

	return nil
}

// NewFilePresenter is constructor of filePresenter

// If path for output file is empty, then output file will be default

func NewFilePresenter(outputFile string) *FilePresenter {
	if outputFile == "" {
		outputFile = defaultPath
	}

	return &FilePresenter{outputFile: outputFile}
}
