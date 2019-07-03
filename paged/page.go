package paged

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type result struct {
	Count    int
	Next     string
	Previous string
	Results  json.RawMessage
}

type Resource struct {
	Base   string
	page   *pagedResult
	Client *http.Client
}

func (p *Resource) More() bool {
	if p.page == nil {
		p.page = &pagedResult{
			Next: p.Base,
		}
	}

	if p.page.Next == "" {
		return false
	}

	n := p.page.Next

	resp, err := p.Client.Get(n)
	if err != nil {
		log.Println(err)
		return false
	}

	pp := &result{}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false
	}

	if err := json.Unmarshal(b, pp); err != nil {
		log.Println(err)
		return false
	}

	p.page = pp

	return true
}

func (p *Resource) Results() []byte {
	return p.page.Results
}
