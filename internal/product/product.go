package product

import "example/model"

type ProductResp struct {
	Records   []model.Product `json:"records"`
	Page      int             `json:"page"`
	PageSize  int             `json:"page_size"`
	TotalPage int             `json:"total_page"`
}
