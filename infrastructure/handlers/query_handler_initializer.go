package handlers

import (
	"presentation-advert-read-api/application/handlers"
	"presentation-advert-read-api/application/repository"
	"presentation-advert-read-api/application/tracers"
	"presentation-advert-read-api/infrastructure/handlers/query_handlers"
	infraTracers "presentation-advert-read-api/infrastructure/tracers"
)

func InitializeQueryHandler(
	categoryRepository repository.CategoryRepository,
	advertRepository repository.AdvertRepository,
) (*handlers.QueryHandler, error) {
	tracer := []tracers.Tracer{
		infraTracers.NewExampleTracer(),
	}
	commandHandler := &handlers.QueryHandler{}
	commandHandler.GetAdvert = handlers.NewQueryHandlerDecorator(query_handlers.NewGetAdvertQueryHandler(
		advertRepository,
	), tracer)
	commandHandler.GetCategory = handlers.NewQueryHandlerDecorator(query_handlers.NewGetCategoryQueryHandler(
		categoryRepository),
		tracer)
	return commandHandler, nil
}
