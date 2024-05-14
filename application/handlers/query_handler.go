package handlers

import (
	"presentation-advert-read-api/application/queries"
	"presentation-advert-read-api/model/model_api"
)

type QueryHandler struct {
	GetAdvert   QueryHandlerDecorator[*queries.GetAdvertQuery, *model_api.AdvertResponse]
	GetCategory QueryHandlerDecorator[*queries.GetCategoryQuery, *model_api.CategoryResponse]
}
