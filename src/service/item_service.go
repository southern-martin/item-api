package service

import (
	"github.com/southern-martin/item-api/src/domain/item"
	"github.com/southern-martin/item-api/src/domain/query"
	"github.com/southern-martin/util-go/rest_error"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(item.Item) (*item.Item, rest_error.RestErr)
	Get(string) (*item.Item, rest_error.RestErr)
	Search(query.EsQuery) ([]item.Item, rest_error.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item item.Item) (*item.Item, rest_error.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*item.Item, rest_error.RestErr) {
	item := item.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(query query.EsQuery) ([]item.Item, rest_error.RestErr) {
	dao := item.Item{}
	return dao.Search(query)
}
