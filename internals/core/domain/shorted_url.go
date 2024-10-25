package domain

type ShortedUrl struct {
	Id          string `json:"id"`
	OriginalUrl string `json:"original_url"`
	Available   bool   `json:"available"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
