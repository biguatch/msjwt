package msjwt

type Config struct {
	Audience string `json:"audience"`
	Issuer   string `json:"issuer"`
	Secret   string `json:"secret"`
}

func (c *Config) GetAudience() string {
	return c.Audience
}

func (c *Config) GetIssuer() string {
	return c.Issuer
}

func (c *Config) GetSecret() string {
	return c.Secret
}
