package smart_solution

import "example/model"

type SmartSolutionResp struct {
	Records   []model.SmartSolution `json:"records"`
	Page      int                   `json:"page"`
	PageSize  int                   `json:"page_size"`
	TotalPage int                   `json:"total_page"`
}
