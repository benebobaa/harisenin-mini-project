package request

import "github.com/benebobaa/harisenin-mini-project/domain/entity"

type RequestQuoteDTO struct {
	Text string `json:"text"`
}

func (this RequestQuoteDTO) ToQuoteEntity() entity.Quote {
	return entity.Quote{
		Text: this.Text,
	}
}
