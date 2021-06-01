package item

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/southern-martin/item-api/src/client/elasticsearch"
	"github.com/southern-martin/item-api/src/domain/query"
	"github.com/southern-martin/item-api/src/domain/query"
	"github.com/southern-martin/util-go/rest_error"
	"strings"
)

const (
	indexItems = "item"
	typeItem   = "_doc"
)

func (i *Item) Save() rest_error.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_error.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_error.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_error.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
		}
		return rest_error.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}

	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_error.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}
	i.Id = itemId
	return nil
}

func (i *Item) Search(query query.EsQuery) ([]Item, rest_error.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, rest_error.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_error.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		item.Id = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, rest_error.NewNotFoundError("no item found matching given criteria")
	}
	return items, nil
}
