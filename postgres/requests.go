package postgres

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/jmagoga/new-equimper-go-graphql/graph/model"
)

type RequestsRepo struct {
	DB *pg.DB
}

func (r * RequestsRepo) GetRequests(filter *model.RequestFilter, limit, offset *int) ([]*model.Request, error) {
	var requests []*model.Request

	query := r.DB.Model(&requests).Order("id")

	if filter != nil {
		if filter.FullText != nil || *filter.FullText != "" {
			query.Where("full_text ILIKE ?", fmt.Sprintf("%%%s%%", *filter.FullText))
		}
	}

	if limit != nil {
		query.Limit(*limit)
	}

	if offset != nil {
		query.Offset(*offset)
	}

	err := query.Select()
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *RequestsRepo) CreateRequest(request *model.Request /* model.NewRequest ??? */) (*model.Request, error) {
	_, err := r.DB.Model(request).Returning("*").Insert()

	return request, err
}