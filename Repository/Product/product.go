package Product

import (
	"akasia/Config"
	"akasia/Controller/Dto/Request"
	"akasia/Controller/Dto/Response"
	"context"
	"time"
)

func (p *product) CreateProduct(ctx context.Context, param Request.CreateProduct) (err error) {
	query := `INSERT INTO t_product (id, title, description, rating, image) VALUES (?, ?, ?, ?, ?)`
	if _, err = Config.DATABASE_MAIN.Get().ExecContext(ctx, query, param.Id, param.Title, param.Description, param.Rating, param.Image); err != nil {
		return err
	}

	return
}

func (p *product) CheckExistsProductTitle(ctx context.Context, title string) (exists bool, err error) {
	query := `SELECT EXISTS (SELECT 1 FROM t_product WHERE LOWER(title) = LOWER(?))`
	err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, title).Scan(&exists)
	return
}

func (p *product) CheckExistsProductId(ctx context.Context, id string) (exists bool, err error) {
	query := `SELECT EXISTS (SELECT 1 FROM t_product WHERE id = ?)`
	err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, id).Scan(&exists)
	return
}

func (p *product) UpdateProduct(ctx context.Context, param Request.UpdateProduct) (err error) {
	query := `UPDATE t_product SET title = ?, description = ?, rating = ?, image = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL`
	if _, err = Config.DATABASE_MAIN.Get().ExecContext(ctx, query, param.Title, param.Description, param.Rating, param.Image, time.Now(), param.Id); err != nil {
		return err
	}
	return
}

func (p *product) DeleteProduct(ctx context.Context, id string) (err error) {
	query := `UPDATE t_product SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL `
	if _, err = Config.DATABASE_MAIN.Get().ExecContext(ctx, query, time.Now(), id); err != nil {
		return err
	}

	return
}

func (p *product) ListProduct(ctx context.Context, sortBy string) (res []Response.ProductList, err error) {
	var (
		data       Response.ProductList
		connection = Config.DATABASE_MAIN.Get()
	)

	query := `SELECT id, title, description, rating, image FROM t_product WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT 10`
	if sortBy == "title" {
		query = `SELECT id, title, description, rating, image FROM t_product WHERE deleted_at IS NULL ORDER BY title ASC LIMIT 10`
	}

	if sortBy == "rating" {
		query = `SELECT id, title, description, rating, image FROM t_product WHERE deleted_at IS NULL ORDER BY rating ASC LIMIT 10`
	}

	rows, err := connection.QueryContext(ctx, query)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&data.Id, &data.Title, &data.Description, &data.Rating, &data.Image); err != nil {
			return
		}

		res = append(res, data)
	}

	return
}

func (p *product) DetailProduct(ctx context.Context, id string) (res Response.ProductDetail, err error) {
	query := `SELECT id, title, description, rating, image, updated_at FROM t_product WHERE id = ? AND deleted_at IS NULL`
	if err = Config.DATABASE_MAIN.Get().QueryRowContext(ctx, query, id).Scan(&res.Id, &res.Title, &res.Description,
		&res.Rating, &res.Image, &res.UpdatedAt); err != nil {
		return
	}
	return
}
