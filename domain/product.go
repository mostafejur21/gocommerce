package domain

// main business logic will be inside this domain
// model or entity -> existence
type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImageUrl    string  `json:"imageUrl" db:"img_url"`
}
