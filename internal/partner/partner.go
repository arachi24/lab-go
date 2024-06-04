package partner

type UpdateReq struct {
	Image       string `json:"image"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
