package main

import (
	"os"

	service "github.com/pavelk123/SpamMaskerV2/service"
	slog "golang.org/x/exp/slog"
)

func main() {
	//"./test/e.txt"
	//"./test/output.txt"
	var inputFile string
	var outputFile string

	paths := os.Args[1:]

	switch len(paths) {
	case 0:
		slog.Error("Params not provided")
		return
	case 1:
		inputFile = paths[0]
		outputFile = ""
	case 2:
		inputFile = paths[0]
		outputFile = paths[1]
	default:
		slog.Error("Toooooo much params")
		return
	}

	fileProducer := service.NewFileProducer(inputFile)
	filePresenter := service.NewFilePresenter(outputFile)

	service := service.NewService(fileProducer, filePresenter)
	err := service.Run()

	if err != nil {
		slog.Error(err.Error())
		return
	}

	slog.Info("Complete")
}
