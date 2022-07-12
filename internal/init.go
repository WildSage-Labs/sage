package internal

func NewSage() (error, *Sage) {
	return nil, &Sage{}
}

func (s *Sage) init() error {
	return nil
}

func (s *Sage) Start() {}
