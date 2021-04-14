package api

type LoadingMessageResponse struct {
	BaseGame string `json:"base_game,omitempty"`
	SubGame  string `json:"sub_game,omitempty"`
	Message  string `json:"message"`
}
