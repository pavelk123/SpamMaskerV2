package main

import (
	"fmt"

	"./service"
)

func main() {
	fileProducer := service.NewFileProducer()
	filePresenter := service.NewFilePresenter("./test/output.txt")

	service := service.NewService(fileProducer, filePresenter)

	err := service.Run("./test/e.txt", maskingUrl)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Complete")
}

func maskingUrl(str string) string {
	var startUrlIndex, isMasking = 0, false
	buffer := []byte(str)

	for i := range buffer {
		if buffer[i] == 'h' && string(buffer[i:i+7]) == "http://" {
			startUrlIndex = i + 7
			isMasking = true
		}

		if startUrlIndex != 0 && i >= startUrlIndex && isMasking {
			if buffer[i] == ' ' {
				isMasking = false
			} else {
				buffer[i] = '*'
			}
		}
	}
	return string(buffer)
}
