package director

type UpdateReq struct {
	Avatar      string `json:"avatar"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Position    string `json:"position"`
	Description string `json:"description"`
}
