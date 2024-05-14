package controller

import (
	"github.com/labstack/echo/v4"
	"presentation-advert-read-api/application/handlers"
	"presentation-advert-read-api/application/queries"
	"presentation-advert-read-api/infrastructure/configuration/custom_error"
	"strconv"
)

type categoryController struct {
	queryHandler *handlers.QueryHandler
}

func NewCategoryController(
	echo *echo.Echo,
	queryHandler *handlers.QueryHandler,
) {
	controller := &categoryController{
		queryHandler: queryHandler,
	}
	controller.register(echo)
}

func (controller *categoryController) register(e *echo.Echo) {
	e.GET("/categories/:id", controller.GetCategoryById)

}

// GetCategoryById godoc
// @tags categories
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success  200  {object}  model_api.CategoryResponse
// @Failure  400  {object} custom_error.CustomError
// @Failure  404  {object} custom_error.CustomError
// @Router /categories/{id} [get]
func (controller *categoryController) GetCategoryById(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return custom_error.BadRequestErr("id must be number")
	}
	categoryResponse, err := controller.queryHandler.GetCategory.Handle(ctx, &queries.GetCategoryQuery{Id: id})
	if err != nil {
		return err
	}
	return c.JSON(200, categoryResponse)
}
