package productarea

import "example/model"

type UpdateReq struct {
	Banner      *model.Banner `json:"banner"`
	Logo        string        `json:"logo"`
	Image       string        `json:"image"`
	Link        string        `json:"link"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
}
