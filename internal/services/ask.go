package services

type AskService interface {
	Process() (string, error)
}

type ask struct {
	query string
}

func NewAskService(query string) *ask {
	return &ask{query}
}

func (a *ask) Process() (string, error) {
	return a.query, nil
}
