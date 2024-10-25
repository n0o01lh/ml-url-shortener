package domain

type ShortRequest struct {
	Url       string `json:"url" validate:"required"`
	Available bool   `json:"available"`
}

func NewShortRequest(url string, available bool) *ShortRequest {
	return &ShortRequest{
		Url:       url,
		Available: available,
	}
}
