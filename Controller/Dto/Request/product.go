package Request

type CreateProduct struct {
	Id          string  `json:"id"`
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float64 `json:"rating" validate:"gte=1,lte=5"`
	Image       string  `json:"image"`
}

type UpdateProduct struct {
	Id          string  `json:"id"`
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float64 `json:"rating" validate:"gte=1,lte=5"`
	Image       string  `json:"image"`
}
