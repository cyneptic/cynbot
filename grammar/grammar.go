package grammar

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type GrammarInterface interface {
	Check(string) (string, string, string, error)
}

type SomeStruct struct {
}

type GrammarChecker struct {
	Sentence          string
	CorrectedSentence string
	Changes           []string
	APIAddress        string
}

func NewGrammarChecker(sentence string) *GrammarChecker {
	uid, token := "FyJqu7op3asCFv0D", "12050"
	println(uid, token)
	address := fmt.Sprintf("https://www.stands4.com/services/v2/grammar.php?uid=%s&tokenid=%s&text=%s&format=json", uid, token, url.QueryEscape(sentence))

	return &GrammarChecker{
		Sentence:   sentence,
		APIAddress: address,
	}
}

func (g *GrammarChecker) Check(initialSentence string) (string, string, string, error) {
	cl := &http.Client{
		Transport: &http.Transport{},
	}

	resp, err := cl.Get(g.APIAddress)
	if err != nil {
		log.Printf("error doing grammar check when sending request, err : %v", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error doing grammar check when checking response, err : %v", err)
	}

	fmt.Println(string(b))

	return initialSentence, initialSentence, "", nil
}
