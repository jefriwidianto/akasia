package Response

type ProductList struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      string `json:"rating"`
	Image       string `json:"image"`
}

type ProductDetail struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      string  `json:"rating"`
	Image       string  `json:"image"`
	UpdatedAt   *string `json:"updatedAt"`
}
