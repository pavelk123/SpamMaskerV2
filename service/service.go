package service

type Producer interface {
	Produce( string) ([]string, error)
}

type Presenter interface {
	Present( []string) error
}

type Service struct {
	prod Producer
	pres Presenter
}

func NewService(prod Producer, pres Presenter) *Service {
	return &Service{
		prod: prod,
		pres: pres,
	}
}

func (s *Service) Run(inputFile string,action func(string)string) error {
	
	data,errRead:=s.prod.Produce(inputFile)
	if errRead!= nil{
		return errRead
	}

	for i := range data {
		data[i] = action(data[i])
	}

	errPres :=s.pres.Present(data)
	if errPres != nil{
		return errPres
	}

	return nil
}
