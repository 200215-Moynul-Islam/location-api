package dtos

type LocationResponse struct {
	ID             int    `json:"id"`
	Country        string `json:"country"`
	State          string `json:"state"`
	City           string `json:"city"`
	BaseImageURL   string `json:"baseImageUrl"`
}

type LocationImageResponse struct {
	ID       int    `json:"id"`
	ImageURL string `json:"imageUrl"`
}

type UpdateBaseImageRequest struct {
	LocationImageID int `json:"locationImageId"`
}
