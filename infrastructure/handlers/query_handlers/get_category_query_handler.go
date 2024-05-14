package query_handlers

import (
	"context"
	"presentation-advert-read-api/application/handlers"
	"presentation-advert-read-api/application/queries"
	"presentation-advert-read-api/application/repository"
	"presentation-advert-read-api/model/model_api"
)

type getCategoryQueryHandler struct {
	categoryRepository repository.CategoryRepository
}

func NewGetCategoryQueryHandler(
	categoryRepository repository.CategoryRepository,
) handlers.QueryHandlerInterface[*queries.GetCategoryQuery, *model_api.CategoryResponse] {
	return &getCategoryQueryHandler{
		categoryRepository: categoryRepository,
	}
}

func (handler *getCategoryQueryHandler) Handle(ctx context.Context, query *queries.GetCategoryQuery) (*model_api.CategoryResponse, error) {
	category, err := handler.categoryRepository.GetById(ctx, query.Id)
	if err != nil {
		return nil, err
	}
	return &model_api.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}, nil
}
