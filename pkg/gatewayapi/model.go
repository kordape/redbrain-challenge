package gatewayapi

type HashMovieNameRequest struct {
	Movie string `json:"movie"`
}

type HashMovieNameResponse struct {
	Encrypted string `json:"encrypted,omitempty"`
	Error     string `json:"error,omitempty"`
}
