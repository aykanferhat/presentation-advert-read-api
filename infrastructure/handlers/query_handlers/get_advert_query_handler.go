package query_handlers

import (
	"context"
	"presentation-advert-read-api/application/handlers"
	"presentation-advert-read-api/application/queries"
	"presentation-advert-read-api/application/repository"
	"presentation-advert-read-api/model/model_api"
)

type getAdvertQueryHandler struct {
	advertRepository repository.AdvertRepository
}

func NewGetAdvertQueryHandler(
	advertRepository repository.AdvertRepository,
) handlers.QueryHandlerInterface[*queries.GetAdvertQuery, *model_api.AdvertResponse] {
	return &getAdvertQueryHandler{
		advertRepository: advertRepository,
	}
}

func (handler *getAdvertQueryHandler) Handle(ctx context.Context, query *queries.GetAdvertQuery) (*model_api.AdvertResponse, error) {
	advert, err := handler.advertRepository.GetById(ctx, query.Id)
	if err != nil {
		return nil, err
	}
	return &model_api.AdvertResponse{
		Id:          advert.Id,
		Title:       advert.Title,
		Description: advert.Description,
		Category: model_api.AdvertCategoryResponse{
			Id:   advert.Category.Id,
			Name: advert.Category.Name,
		},
	}, nil
}
