package gogs

type WebauthnContainer struct {
	WebauthnData string `json:"webauthn_data"`
}

func (c *Client) IsWebauthnEnabled() (*bool, error) {
	var enabled bool
	return &enabled, c.getParsedResponse("GET", "/webauthn/isenabled", nil, nil, &enabled)
}
