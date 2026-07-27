package dtos

type LocationResponse struct {
	ID             int    `json:"id"`
	Country        string `json:"country"`
	State          string `json:"state"`
	City           string `json:"city"`
	BaseImageURL   string `json:"baseImageUrl"`
}
