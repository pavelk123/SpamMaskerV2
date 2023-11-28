package service

import (
	"fmt"
)

type producer interface {
	produce() ([]string, error)
}

type presenter interface {
	present([]string) error
}

//Service is structure for masking url service
//Including inside 2 fields:
//producer - for data provider unit
//presenter - for data presenter unit
type Service struct {
	prod producer
	pres presenter
}

//Run is method for start Service working
func (s *Service) Run() error {
	errPrefix := "Service.Run:"

	data, err := s.prod.produce()
	if err != nil {
		return fmt.Errorf("%s %w", errPrefix, err)
	}

	for i := range data {
		data[i] = s.maskingUrl(data[i])
	}

	if err = s.pres.present(data); err != nil {
		return fmt.Errorf("%s %w", errPrefix, err)
	}

	return nil
}

func (s Service) maskingUrl(str string) string {
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
				continue
			}

			buffer[i] = '*'
		}
	}
	return string(buffer)
}

//NewService - constructor for Service for masking Urls spam
func NewService(prod producer, pres presenter) *Service {
	return &Service{
		prod: prod,
		pres: pres,
	}
}
