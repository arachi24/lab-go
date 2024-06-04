package homepage

import "example/model"

type UpdateReq struct {
	MainBanner  *model.Banner `json:"main_banner"`
	SubBanner   *model.Banner `json:"sub_banner"`
	Title       string        `json:"title"`
	SubTitle    string        `json:"sub_title"`
	Description string        `json:"description"`
}
