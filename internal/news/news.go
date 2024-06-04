package news

type UpdateReq struct {
	Title       string `json:"title"`
	SubTitle    string `json:"sub_title"`
	Description string `json:"description"`
}
