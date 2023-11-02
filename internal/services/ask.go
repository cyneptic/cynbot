package services

import (
	"fmt"

	"github.com/cyneptic/cynbot/utils"
)

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
	res, err := utils.GetWholeText(a.query)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
	return res, err
}
