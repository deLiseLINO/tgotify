package telegram

type UpdatesResponse struct {
	Ok      bool     `json:"ok"`
	Updates []Update `json:"result"`
}

type Update struct {
	ID      int    `json:"update_id"`
	Message string `json:"message"`
}
