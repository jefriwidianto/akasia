package Product

import (
	"akasia/Controller/Dto/Request"
	"akasia/Controller/Dto/Response"
	"context"
)

type RepositoryProduct interface {
	CreateProduct(ctx context.Context, param Request.CreateProduct) (err error)
	CheckExistsProductTitle(ctx context.Context, title string) (exists bool, err error)
	CheckExistsProductId(ctx context.Context, id string) (exists bool, err error)
	UpdateProduct(ctx context.Context, param Request.UpdateProduct) (err error)
	DeleteProduct(ctx context.Context, id string) (err error)
	ListProduct(ctx context.Context, sortBy string) (res []Response.ProductList, err error)
	DetailProduct(ctx context.Context, id string) (res Response.ProductDetail, err error)
}

type product struct{}

func NewRepository() RepositoryProduct {
	return &product{}
}
