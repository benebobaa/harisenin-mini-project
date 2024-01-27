package usecase

import (
	"github.com/benebobaa/harisenin-mini-project/domain/entity"
	"github.com/benebobaa/harisenin-mini-project/domain/repository"
)

type QuoteUsecase interface {
	GetRandomQuote() (*entity.Quote, error)
	SaveQuote(action entity.Quote) error
}

type quoteUsecaseImpl struct {
	databaseRepository repository.DatabaseRepository
}

func NewQuoteUsecase(databaseRepository repository.DatabaseRepository) quoteUsecaseImpl {
	return quoteUsecaseImpl{
		databaseRepository: databaseRepository,
	}
}

func (this *quoteUsecaseImpl) GetRandomQuote() (*entity.Quote, error) {
	var quote entity.Quote
	if err := this.databaseRepository.FindRandomQuote(&quote); err != nil {
		return nil, err
	}

	return &quote, nil
}

func (this *quoteUsecaseImpl) SaveQuote(action entity.Quote) error {
	if err := this.databaseRepository.Create(&action); err != nil {
		return err
	}

	return nil
}
