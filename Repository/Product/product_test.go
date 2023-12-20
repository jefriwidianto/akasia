package Product

import (
	"akasia/Config"
	"akasia/Controller/Dto/Request"
	"akasia/Controller/Dto/Response"
	"context"
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func ConnectionMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println(err)
	}

	return db, mock
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestCreateProductSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = Request.CreateProduct{
		Id:          uuid.NewV4().String(),
		Title:       "Title",
		Description: "Desc",
		Rating:      2.1,
		Image:       "product.jpg",
	}

	query := `INSERT INTO t_product (id, title, description, rating, image) VALUES (?, ?, ?, ?, ?)`
	mock.ExpectExec(query).WithArgs(paramsMock.Id, paramsMock.Title, paramsMock.Description, paramsMock.Rating, paramsMock.Image).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().CreateProduct(context.Background(), paramsMock)
	assert.NoError(t, err)
}

func TestCreateProductFailure(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = Request.CreateProduct{
		Id:          uuid.NewV4().String(),
		Title:       "Title",
		Description: "Desc",
		Rating:      2,
		Image:       "product.jpg",
	}

	query := `INSERT INTO t_product (title, description, rating, image) VALUES (?, ?, ?, ?)`
	mock.ExpectExec(query).WithArgs(paramsMock.Title, paramsMock.Description, paramsMock.Rating, paramsMock.Image).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().CreateProduct(context.Background(), paramsMock)
	assert.Error(t, err)
}

func TestListProductSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var resMock = []Response.ProductList{
		Response.ProductList{
			Id:          "0e29c9ab-5a07-4134-98c6-c834a06891c2",
			Title:       "Baju Anak",
			Description: "baju anak",
			Rating:      5,
			Image:       "baju.jpg",
		},
	}

	var paramMock = "title"
	query := `SELECT id, title, description, rating, image FROM t_product WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT 10`
	if paramMock == "title" {
		query = `SELECT id, title, description, rating, image FROM t_product WHERE deleted_at IS NULL ORDER BY title ASC LIMIT 10`
	}

	rowsTitle := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image"}).
		AddRow(resMock[0].Id, resMock[0].Title, resMock[0].Description, resMock[0].Rating, resMock[0].Image)

	mock.ExpectQuery(query).WithArgs().WillReturnRows(rowsTitle)

	resTitle, err := NewRepository().ListProduct(context.Background(), paramMock)
	assert.Nil(t, err)
	assert.Equal(t, resMock, resTitle)

	paramMock = "rating"
	if paramMock == "rating" {
		query = `SELECT id, title, description, rating, image FROM t_product WHERE deleted_at IS NULL ORDER BY rating ASC LIMIT 10`
	}

	rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image"}).
		AddRow(resMock[0].Id, resMock[0].Title, resMock[0].Description, resMock[0].Rating, resMock[0].Image)

	mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)

	res, err := NewRepository().ListProduct(context.Background(), paramMock)
	assert.Nil(t, err)
	assert.Equal(t, resMock, res)
}

func TestCheckExistsProductTitleSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramMock = "Baju Anak"
	var resMock bool

	query := `SELECT EXISTS (SELECT 1 FROM t_product WHERE LOWER(title) = LOWER(?)) AS "exists"`

	rows := sqlmock.NewRows([]string{"exists"}).AddRow(resMock)
	mock.ExpectQuery(query).WithArgs(paramMock).WillReturnRows(rows)

	res, err := NewRepository().CheckExistsProductTitle(context.Background(), paramMock)
	assert.NoError(t, err)
	assert.Equal(t, resMock, res)

	paramMockOtherValue := "Lukisan"
	query = `SELECT EXISTS (SELECT 1 FROM t_product WHERE LOWER(title) = LOWER(?)) AS "exists"`
	rows = sqlmock.NewRows([]string{"exists"}).AddRow(resMock)
	mock.ExpectQuery(query).WithArgs(paramMockOtherValue).WillReturnRows(rows)

	resOtherValue, err := NewRepository().CheckExistsProductTitle(context.Background(), paramMockOtherValue)
	assert.NoError(t, err)
	assert.Equal(t, resMock, resOtherValue)
}

func TestCheckExistsProductIdSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramMock = "1a41c84f-9aba-4b77-acee-9dae4fc06b4a"
	var resMock bool

	query := `SELECT EXISTS (SELECT 1 FROM t_product WHERE id = ?) AS "exists"`

	rows := sqlmock.NewRows([]string{"exists"}).AddRow(resMock)
	mock.ExpectQuery(query).WithArgs(paramMock).WillReturnRows(rows)

	res, err := NewRepository().CheckExistsProductId(context.Background(), paramMock)
	assert.NoError(t, err)
	assert.Equal(t, resMock, res)

	paramMockOtherValue := "0e29c9ab-5a07-4134-98c6-c834a06891c2"
	query = `SELECT EXISTS (SELECT 1 FROM t_product WHERE id = ?) AS "exists"`
	rows = sqlmock.NewRows([]string{"exists"}).AddRow(resMock)
	mock.ExpectQuery(query).WithArgs(paramMockOtherValue).WillReturnRows(rows)

	resOtherValue, err := NewRepository().CheckExistsProductId(context.Background(), paramMockOtherValue)
	assert.NoError(t, err)
	assert.Equal(t, resMock, resOtherValue)
}

func TestUpdatedProductSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = Request.UpdateProduct{
		Id:          "0e29c9ab-5a07-4134-98c6-c834a06891c2",
		Title:       "Title",
		Description: "Desc",
		Rating:      2.1,
		Image:       "product.jpg",
	}

	query := `UPDATE t_product SET title = ?, description = ?, rating = ?, image = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL`
	mock.ExpectExec(query).WithArgs(paramsMock.Title, paramsMock.Description, paramsMock.Rating, paramsMock.Image, AnyTime{}, paramsMock.Id).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().UpdateProduct(context.Background(), paramsMock)
	assert.NoError(t, err)
}

func TestUpdateProductFailure(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = Request.UpdateProduct{
		Id:          "0e29c9ab-5a07-4134-98c6-c834a06891c2",
		Title:       "Title",
		Description: "Desc",
		Rating:      2.1,
		Image:       "product.jpg",
	}

	query := `UPDATE t_product SET title = ?, description = ?, rating = ?, image = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL`
	mock.ExpectExec(query).WithArgs(paramsMock.Title, paramsMock.Description, paramsMock.Rating, paramsMock.Image, AnyTime{}).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().UpdateProduct(context.Background(), paramsMock)
	assert.Error(t, err)
}

func TestDeleteProductSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = "0e29c9ab-5a07-4134-98c6-c834a06891c2"

	query := `UPDATE t_product SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL`
	mock.ExpectExec(query).WithArgs(AnyTime{}, paramsMock).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().DeleteProduct(context.Background(), paramsMock)
	assert.NoError(t, err)
}

func TestDeleteProductFailure(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var paramsMock = "1a41c84f-9aba-4b77-acee-9dae4fc06b4ax"

	query := `UPDATE t_product SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL`
	mock.ExpectExec(query).WithArgs(paramsMock).WillReturnResult(
		sqlmock.NewResult(0, 1))

	err := NewRepository().DeleteProduct(context.Background(), paramsMock)
	assert.Error(t, err)
}

func TestDetailProductIdSuccess(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	var updatedAtValue = "2023-12-19 16:47:51"
	var paramMock = "1a41c84f-9aba-4b77-acee-9dae4fc06b4a"
	var resMock = Response.ProductDetail{
		Id:          "1a41c84f-9aba-4b77-acee-9dae4fc06b4a",
		Title:       "Baju Dewasa",
		Description: "baju anak",
		Rating:      2,
		Image:       "aju_dewasa.jpg",
		UpdatedAt:   &updatedAtValue,
	}

	query := `SELECT id, title, description, rating, image, updated_at FROM t_product WHERE id = ? AND deleted_at IS NULL`

	rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "updated_at"}).AddRow(resMock.Id,
		resMock.Title, resMock.Description, resMock.Rating, resMock.Image, resMock.UpdatedAt)
	mock.ExpectQuery(query).WithArgs(paramMock).WillReturnRows(rows)

	res, err := NewRepository().DetailProduct(context.Background(), paramMock)
	assert.NoError(t, err)
	assert.Equal(t, resMock, res)

	paramMockOtherValue := "0e29c9ab-5a07-4134-98c6-c834a06891c2"
	var resMockOtherValue = Response.ProductDetail{
		Id:          "0e29c9ab-5a07-4134-98c6-c834a06891c2",
		Title:       "Baju Anak",
		Description: "baju anak",
		Rating:      5,
		Image:       "baju.jpg",
		UpdatedAt:   &updatedAtValue,
	}

	query = `SELECT id, title, description, rating, image, updated_at FROM t_product WHERE id = ? AND deleted_at IS NULL`
	rows = sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "updated_at"}).AddRow(resMock.Id,
		resMock.Title, resMock.Description, resMock.Rating, resMock.Image, resMock.UpdatedAt)
	mock.ExpectQuery(query).WithArgs(paramMockOtherValue).WillReturnRows(rows)

	resOtherValue, err := NewRepository().DetailProduct(context.Background(), paramMockOtherValue)
	assert.NoError(t, err)
	assert.NotEqual(t, resMockOtherValue, resOtherValue)
}
