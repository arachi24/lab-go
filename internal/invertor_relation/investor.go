package invertorrelation

import "example/model"

type UpdateReq struct {
	Banner      []*model.Banner `json:"banner"`
	Image       string          `json:"image"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
}
