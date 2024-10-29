package domain

type Stats struct {
	Id     string `json:"id"`
	Clicks int64  `json:"clicks"`
}

func NewStats(id string, clicks int64) *Stats {
	return &Stats{
		Id:     id,
		Clicks: clicks,
	}
}
