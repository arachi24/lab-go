package business

type UpdateReq struct {
	Logo        string `json:"logo"`
	Image       string `json:"image"`
	Link        string `json:"link"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
