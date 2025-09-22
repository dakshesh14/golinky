package link

type ShortenURLRequest struct {
	URL string `json:"uri" validate:"required,min=10"`
}
