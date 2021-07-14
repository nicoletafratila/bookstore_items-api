package items

import (
	"errors"
	"github.com/nicoletafratila/bookstore_items-api/clients/elasticsearch"
	"github.com/nicoletafratila/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, 1)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}

	i.Id = result.Id
	return nil
}
