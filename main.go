package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DataProducer interface {
	Produce(data string) ([]string, error)
}
type ResultProcessor interface {
	Process(data []string) error
}

type Service interface {
	DataProducer
	ResultProcessor
}

type StringProvider struct{}

func (sp StringProvider) Produce(data string) ([]string, error) {
	return []string{strings.TrimSpace(data)}, nil
}

func (sp StringProvider) Process(data []string) error {
	text := strings.Join(data, " ")
	fmt.Println(text)

	return nil
}

type FileProvider struct {
	filePath string
}

func (fp *FileProvider) Produce(data string) ([]string, error) {
	//./e.txt
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

func (fp FileProvider) Process(data []string) error {
	file, err := os.Create(fp.filePath)
	writer := bufio.NewWriter(file)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, str := range data {
		_,err = writer.WriteString(str)
		if err!=nil{
			return err
		}
	}
	
	writer.Flush()
	return nil
}

func main() {

	const (
		StringInput = iota + 1
		FileInput
	)

	var service Service
	var input string

	fmt.Printf("Select data input method (%d - string, %d - file): ", StringInput, FileInput)
	fmt.Scanf("%s\n", &input)

	switch input {
	case strconv.Itoa(StringInput):
		service = StringProvider{}
		fmt.Print("Input str: ")

	case strconv.Itoa(FileInput):
		service = &FileProvider{}
		fmt.Print("Input path: ")

	default:
		fmt.Println("Incorrect input")
		return
	}

	input, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	buffer, err := service.Produce(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range buffer {
		buffer[i] = MaskingUrl(buffer[i])
	}

	err = service.Process(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Complete")
}

func MaskingUrl(str string) string {
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
