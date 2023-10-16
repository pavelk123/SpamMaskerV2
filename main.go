package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	pc "./processor"
	pd "./producer"
)

type Service struct {
	Producer  pd.Producer
	Processor pc.Processor
}

func main() {
	const (
		StringInput = iota + 1
		FileInput
	)

	var service Service = Service{}
	var input string
	var errInput error

	fmt.Printf("Select data input method (%d - string, %d - file): ", StringInput, FileInput)
	fmt.Scanf("%s\n", &input)

	switch input {
	case strconv.Itoa(StringInput):
		service.Producer = pd.StringProducer{}
		fmt.Print("Input str: ")

	//./test/e.txt
	case strconv.Itoa(FileInput):
		service.Producer = &pd.FileProducer{}
		fmt.Print("Input path: ")

	default:
		fmt.Println("Incorrect input")
		return
	}

	input, errInput = bufio.NewReader(os.Stdin).ReadString('\n')
	if errInput != nil {
		fmt.Println(errInput)
		return
	}

	buffer, err := service.Producer.Produce(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range buffer {
		buffer[i] = maskingUrl(buffer[i])
	}

	if fp, ok := service.Producer.(*pd.FileProducer); ok {
		path := fp.GetFilePath()
		processor := pc.FileProcessor{}
		processor.Init(path)
		service.Processor = &processor

	} else if _, ok := service.Producer.(pd.StringProducer); ok {
		service.Processor = pc.StringProcessor{}
	} else {
		fmt.Println("Problem with Processor")
		return
	}

	err = service.Processor.Process(buffer)
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
