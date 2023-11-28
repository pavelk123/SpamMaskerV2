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
	errPrefix := "filePresenter.present:"

	file, err := os.Create(fp.outputFile)
	if err != nil {
		return fmt.Errorf("%s %w", errPrefix, err)
	}

	writer := bufio.NewWriter(file)
	defer func() {
		if errDefer := file.Close(); err != nil {
			err = fmt.Errorf("%s %w", errPrefix, errDefer)
		}
	}()

	for _, str := range data {
		_, err = writer.WriteString(str)
		if err != nil {
			return fmt.Errorf("%s %w", errPrefix, err)
		}
	}

	if err = writer.Flush(); err != nil {
		return fmt.Errorf("%s %w", errPrefix, err)
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
