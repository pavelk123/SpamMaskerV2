package service

type Producer interface {
	Produce() ([]string, error)
}

type Presenter interface {
	Present([]string) error
}

type Service struct {
	prod Producer
	pres Presenter
}

func (s *Service) Run() error {

	data, err := s.prod.Produce()
	if err != nil {
		return err
	}

	for i := range data {
		data[i] = s.maskingUrl(data[i])
	}

	err = s.pres.Present(data)
	if err != nil {
		return err
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
			} else {
				buffer[i] = '*'
			}
		}
	}
	return string(buffer)
}

func NewService(prod Producer, pres Presenter) *Service {
	return &Service{
		prod: prod,
		pres: pres,
	}
}
