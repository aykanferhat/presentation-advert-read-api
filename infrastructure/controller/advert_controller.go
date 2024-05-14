package controller

import (
	"github.com/labstack/echo/v4"
	"presentation-advert-read-api/application/handlers"
	"presentation-advert-read-api/application/queries"
	"presentation-advert-read-api/infrastructure/configuration/custom_error"
	"strconv"
)

type advertController struct {
	queryHandler *handlers.QueryHandler
}

func NewAdvertController(
	echo *echo.Echo,
	queryHandler *handlers.QueryHandler,
) {
	controller := &advertController{
		queryHandler: queryHandler,
	}
	controller.register(echo)
}

func (controller *advertController) register(e *echo.Echo) {
	e.GET("/adverts/:id", controller.GetAdvertById)

}

// GetAdvertById godoc
// @tags adverts
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success  200  {object}  model_api.AdvertResponse
// @Failure  400  {object} custom_error.CustomError
// @Failure  404  {object} custom_error.CustomError
// @Router /adverts/{id} [get]
func (controller *advertController) GetAdvertById(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return custom_error.BadRequestErr("id must be number")
	}
	advertResponse, err := controller.queryHandler.GetAdvert.Handle(ctx, &queries.GetAdvertQuery{Id: id})
	if err != nil {
		return err
	}
	return c.JSON(200, advertResponse)
}
