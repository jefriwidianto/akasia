package Controller

import (
	"akasia/Controller/Dto/Request"
	"akasia/Controller/Dto/Response"
	"akasia/Repository"
	"database/sql"
	"errors"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type ProductInterface interface {
	CreateProduct(ctx echo.Context) (err error)
	UpdateProduct(ctx echo.Context) (err error)
	DeleteProduct(ctx echo.Context) (err error)
	ListProduct(ctx echo.Context) (err error)
	DetailProduct(ctx echo.Context) (err error)
}

func (c *Controller) CreateProduct(ctx echo.Context) (err error) {
	var req Request.CreateProduct
	if err = ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	if err = ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	exists, err := Repository.ApplicationRepository.Product.CheckExistsProduct(ctx.Request().Context(), req.Title)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	if exists {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: errors.New("Title Duplicate").Error(),
		})
	}

	req.Id = uuid.NewV4().String()
	if err = Repository.ApplicationRepository.Product.CreateProduct(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response.Responses{
		Data:    "",
		Message: http.StatusText(http.StatusOK),
	})
}

func (c *Controller) UpdateProduct(ctx echo.Context) (err error) {
	var req Request.UpdateProduct
	if err = ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	if err = ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	exists, err := Repository.ApplicationRepository.Product.CheckExistsProduct(ctx.Request().Context(), req.Title)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	if exists {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: errors.New("Title Duplicate").Error(),
		})
	}

	req.Id = ctx.Param("id")
	if err = Repository.ApplicationRepository.Product.UpdateProduct(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response.Responses{
		Data:    "",
		Message: http.StatusText(http.StatusOK),
	})
}

func (c *Controller) DeleteProduct(ctx echo.Context) (err error) {
	productId := ctx.Param("id")
	if err = Repository.ApplicationRepository.Product.DeleteProduct(ctx.Request().Context(), productId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response.Responses{
		Data:    "",
		Message: http.StatusText(http.StatusOK),
	})
}

func (c *Controller) ListProduct(ctx echo.Context) (err error) {
	sortBy := ctx.QueryParams().Get("sortBy")
	list, err := Repository.ApplicationRepository.Product.ListProduct(ctx.Request().Context(), sortBy)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    list,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response.Responses{
		Data:    list,
		Message: http.StatusText(http.StatusOK),
	})
}

func (c *Controller) DetailProduct(ctx echo.Context) (err error) {
	productId := ctx.Param("id")
	detail, err := Repository.ApplicationRepository.Product.DetailProduct(ctx.Request().Context(), productId)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.JSON(http.StatusBadRequest, &Response.Responses{
				Data:    nil,
				Message: errors.New("Data Not Found").Error(),
			})
		}

		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Data:    nil,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response.Responses{
		Data:    detail,
		Message: http.StatusText(http.StatusOK),
	})
}
